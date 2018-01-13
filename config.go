package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type configuration struct {
	Db           string `json:"db"`
	DbConnection string `json:"dbconnection"`
	Port         int16  `json:"port"`
}

func getConfiguration() configuration {
	config, _ := loadConfiguration()

	return config
}

func loadConfiguration() (configuration, error) {
	var config configuration

	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		return config, errors.New("Couldn't read configuration file")
	}
	fmt.Println(string(dat))

	err = json.Unmarshal(dat, &config)
	if err != nil {
		return config, errors.New("Couldn't unmarshal configuration file")
	}

	return config, err
}
