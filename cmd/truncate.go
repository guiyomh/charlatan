package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/azer/logger"
	"github.com/doug-martin/goqu"
	reader "github.com/guiyomh/go-faker-fixtures/internal/reader"

	"github.com/spf13/cobra"
)

var log = logger.New("truncate")

var truncateCmd = &cobra.Command{
	Use:   "truncate",
	Short: "Empty tables related to fixtures",
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
		fixturePath := args[0]
		timer := log.Timer()
		myReader := reader.NewFixtureReader()
		data, err := myReader.Read(fixturePath)

		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)
		myDb, err := sql.Open("mysql", dataSource)
		if err != nil {
			panic(err.Error())
		}
		db := goqu.New("mysql", myDb)

		for tableName := range data {
			log.Info(fmt.Sprintf("Truncate Table : %s", tableName))
			sql, args, _ := db.From(tableName).ToTruncateSql()
			fmt.Println(sql)
			sql = strings.Replace(sql, "\"", "`", -1)
			_, err := db.Exec(sql, args...)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}
		}
		//spew.Dump(tables)
		timer.End("Parse fixture data")
	},
}

func init() {
	rootCmd.AddCommand(truncateCmd)
}
