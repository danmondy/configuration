package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host string       `json:"host"`
	Port string       `json:"port"`
	Db   DbConnection `json:"db"`
}

type DbConnection struct {
	Host     string "json:'host'"
	Port     string "json:'port'"
	Username string "json:'username'"
	Password string "json:'password'"
}

func ReadConfig(file string, config interface{}) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, config)
}
