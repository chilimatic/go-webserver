package main

import (
	"code.google.com/p/gcfg"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

type ConfigFile struct {
	file   string
	Config gcfg
}

type ConfigSetting struct {
	key   string
	value string
}

func (cf *ConfigFile) load(fileName string) (bool, error) {
	body, err := ioutil.ReadFile(fileName)

	if err != nil {
		return false, err
	}

	cf.Content = body

	cf.Config.ReadInto()

	return true, nil
}

func main() {
	cf, err := initConfig()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cf)

}

func initConfig() (*ConfigFile, error) {
	configPath := flag.String("file", "./config/http.conf", "path to the default config")

	if configPath == nil {
		return nil, errors.New("The config path is not empty")
	}

	cf := ConfigFile{file: *configPath}

	return &cf, nil
}
