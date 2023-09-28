package appconfig

import (
	"datagen/types"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Conns        []types.DbConn `json:"conns"`
	TupleSize    int            `json:"tupleSize"`
	BatchSize    int            `json:"batchsize"`
	OutputDir    string         `json:"outputDir"`
	ConfigFormat string         `json:"configFormat"`
}

func GetConf() *AppConfig {

	viper.SetConfigName("datagen") // Register config file name (no extension)
	viper.SetConfigType("json")    // Look for specific type
	viper.AddConfigPath("/etc/datagen")
	home, _ := os.UserHomeDir()
	viper.AddConfigPath(filepath.Join(home, ".datagen"))
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("%v", err)
	}

	conf := &AppConfig{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Infof("unable to decode into config struct, %v", err)
	}

	return conf
}
