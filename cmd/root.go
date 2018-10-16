package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Verbose flag to indicate the verbose mode
	Verbose bool
	// DbUser
	DbUser  string
	DbPass  string
	DbName  string
	DbHost  string
	DbPort  int16
	cfgFile string
	cfgName string = ".charlatan"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s)", cfgName))
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&DbUser, "user", "u", "", "Database username (required)")
	rootCmd.MarkPersistentFlagRequired("user")
	rootCmd.PersistentFlags().StringVarP(&DbPass, "password", "p", "", "Database user password (required)")
	rootCmd.MarkPersistentFlagRequired("pass")
	rootCmd.PersistentFlags().StringVarP(&DbName, "dbname", "d", "", "Database name (required)")
	rootCmd.MarkPersistentFlagRequired("dbname")
	rootCmd.PersistentFlags().StringVarP(&DbHost, "host", "", "127.0.0.1", "Host Database (default is 127.0.0.1)")
	rootCmd.PersistentFlags().Int16VarP(&DbPort, "port", "", 3306, "Database port (default is 3306)")

	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("dbname", rootCmd.PersistentFlags().Lookup("dbname"))
}

func initConfig() {
	hasCfgfile := false
	if cfgFile != "" {
		hasCfgfile = true
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := os.Stat(fmt.Sprintf("%s/%s", home, cfgName)); !os.IsNotExist(err) {
			hasCfgfile = true
			viper.AddConfigPath(home)
			viper.SetConfigName(cfgName)
		}
	}
	if hasCfgfile {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "charlatan",
	Short: "charlatan is a very fast fixtures loaders",
	Long:  `A Fast and Flexible fixtures loader built with love by guiyomh.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
