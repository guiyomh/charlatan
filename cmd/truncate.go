package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/azer/logger"
	"github.com/guiyomh/charlatan/internal/pkg/db"
	"github.com/guiyomh/charlatan/internal/pkg/reader"

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
		dbManagerFactory := db.DbManagerFactory{}
		manager, err := dbManagerFactory.NewDbManager("mysql", DbHost, DbPort, DbUser, DbPass)
		if err != nil {
			panic(err)
		}
		for tableName := range data {
			log.Info(fmt.Sprintf("Truncate Table : %s", tableName))
			_, err := manager.TruncateTable("fixtures", tableName)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}
		}
		timer.End("Parse fixture data")
	},
}

func init() {
	rootCmd.AddCommand(truncateCmd)
}
