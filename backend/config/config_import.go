package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Test_func() string {
	absPath, _ := filepath.Abs("")
	return absPath
}

func Import_yaml_config(config_path string) map[string]DatabaseCredentials {
	yfile, read_data_err := ioutil.ReadFile(config_path)

	if read_data_err != nil {
		log.Fatal(read_data_err)
	}
	data := make(map[string]DatabaseCredentials)

	yaml_err := yaml.Unmarshal(yfile, &data)

	if yaml_err != nil {
		log.Fatal(yaml_err)
	}

	return data
}

type DatabaseCredentials struct {
	Host    string
	Port    string
	User    string
	Passwd  string
	DB      string
	Charset string
}
