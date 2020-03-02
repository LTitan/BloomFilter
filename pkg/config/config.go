package config

import (
	"bytes"
	"io/ioutil"

	"github.com/spf13/viper"
)

func readConfigFile(fileName string) (config []byte, err error) {
	config, err = ioutil.ReadFile(fileName)
	return
}

// Conf .
var Conf *viper.Viper

func init() {
	Reload()
}

// Reload reload conifg
func Reload() {
	Conf = viper.New()
	Conf.SetConfigType("toml")
	config, err := readConfigFile("./config/base.toml")
	if err != nil {
		panic(err)
	}
	if err = Conf.ReadConfig(bytes.NewBuffer(config)); err != nil {
		panic(err)
	}
}
