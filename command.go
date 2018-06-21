package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configure struct {
	BitcoinHost       string
	BitcoinPort       string
	BitcoinUser       string
	BitcoinPass       string
	BitcoinhttpMode   bool
	BitcoinDisableTLS bool
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bitcoin-service",
	Short: "Bitcoin middleware for application",
}

// Execute 命令行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}

func init() {
	config = new(configure)
	config.InitConfig()
}

func (conf *configure) InitConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(HomeDir())
	viper.SetConfigName("bitcoin-service")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using Configure file:", viper.ConfigFileUsed())
	} else {
		log.Fatal("Error: bitcoin.yml not found in: ", HomeDir())
	}

	for key, value := range viper.AllSettings() {
		switch key {
		case "btc_host":
			conf.BitcoinHost = value.(string)
		case "btc_port":
			conf.BitcoinPort = value.(string)
		case "btc_usr":
			conf.BitcoinUser = value.(string)
		case "btc_pass":
			conf.BitcoinPass = value.(string)
		case "btc_http_mode":
			conf.BitcoinhttpMode = value.(bool)
		case "btc_disable_tls":
			conf.BitcoinDisableTLS = value.(bool)
		}
	}
}