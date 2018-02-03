package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type configuration struct {
	Db           string `json:"db"`
	DbConnection string `json:"dbconnection"`
	Port         int16  `json:"port"`
	Debug        bool   `json:"debug"`
}

func getConfiguration() configuration {
	config, _ := loadConfiguration()

	return config
}

func loadConfiguration() (configuration, error) {
	var config configuration

	dat, err := ioutil.ReadFile("data/config.json")
	if err != nil {
		return config, errors.New("Couldn't read configuration file")
	}

	err = json.Unmarshal(dat, &config)
	if err != nil {
		return config, errors.New("Couldn't unmarshal configuration file")
	}

	return config, err
}
