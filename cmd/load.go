package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/azer/logger"
	"github.com/guiyomh/charlatan/internal/app/model"
	"github.com/guiyomh/charlatan/internal/pkg/db"
	"github.com/guiyomh/charlatan/internal/pkg/generator"
	"github.com/guiyomh/charlatan/internal/pkg/reader"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load fixtures from the path",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one arg")
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return fmt.Errorf("the directory '%s' not existing", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.New("cmd")
		timer := log.Timer()
		fixturePath := args[0]
		reader := reader.NewFixtureReader()
		data, err := reader.Read(fixturePath)

		if err != nil {
			log.Error(err.Error())
			panic(err)
		}

		generator := generator.NewGenerator()
		rows, err := generator.GenerateRecords(data)
		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
		rowTree := model.NewTree()
		for _, row := range rows {
			rowTree.Add(row)
		}
		dbManagerFactory := db.DbManagerFactory{}
		manager, err := dbManagerFactory.NewDbManager("mysql", DbHost, DbPort, DbUser, DbPass)
		if err != nil {
			panic(err)
		}
		//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)
		//sqlGenerator := db.NewSqlGenerator("mysql", dataSource)
		rowTree.Walk(func(key string, value *model.Row) {
			if value.HasDependencies() {
				for field, relation := range value.DependencyReference {
					target := rowTree.Find(relation.RecordName)
					if relation.FieldName != "" {
						value.Fields[field] = target.Value.Fields[relation.RecordName]
					} else {
						value.Fields[field] = target.Value.Pk
					}
				}
			}
			sql, params, err := manager.BuildInsertSQL("fixtures", value.TableName, value.Fields)
			if err != nil {
				panic(err)
			}
			result, err := manager.Exec(sql, params)
			if err != nil {
				panic(err)
			}
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}
			value.Pk = lastInsertID
		}, true)

		log.Info(fmt.Sprintf("Nb rows : %d", len(rows)))

		timer.End("Insert record in database")
	},
}
