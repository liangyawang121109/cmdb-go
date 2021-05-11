package configmodel

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Cfg struct {
	Name      string `yaml:name`
	DatabasePort      int    `yaml: databaseport`
	User string
	Passwd string
	DatabaseUrl   string
	Database string
	ServerPort string
}

func ConfigGet() (config Cfg){
	file, err := os.Open("./config.yaml")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	cfg := Cfg{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

