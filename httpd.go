package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

type ConfigFile struct {
	file    string
	Content []byte
}

func (cf *ConfigFile) load(fileName string) (bool, error) {
	var dat map[string]interface{}
	body, err := ioutil.ReadFile(fileName)

	if err != nil {
		return false, err
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return false, err
	}
	fmt.Println(dat["localhost"])
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
	cf.load(cf.file)
	return &cf, nil
}
