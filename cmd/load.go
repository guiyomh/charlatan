package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/azer/logger"

	"github.com/guiyomh/go-faker-fixtures/internal/app/model"
	"github.com/guiyomh/go-faker-fixtures/internal/pkg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load fixtures from the path",
	Long:  `All software has versions. This is go-fixtures's`,
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
		Loader := pkg.NewLoader(args[0])
		Loader.Load()
		fmt.Printf("NB Rows : %d\n", len(Loader.Rows))
		//build binary tree
		rowTree := model.NewTree()
		for _, row := range Loader.Rows {
			rowTree.Add(row)
		}
		timer.End("Build row record with fake data")
		timer = log.Timer()
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)
		//sqlGenerator := pkg.NewSqlGenerator("mysql", "fixtures_user:fixtures_pass@/fixtures?charset=utf8")
		fmt.Println(dataSource)
		sqlGenerator := pkg.NewSqlGenerator("mysql", dataSource)

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
			pk, err := sqlGenerator.ToSQL(value)
			if err != nil {
				panic(err)
			}
			value.Pk = pk
		}, true)

		timer.End("Insert record in database")
	},
}
