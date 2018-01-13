package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

var token = ""
var config configuration
var db *gorm.DB

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Wrong input.")
	} else if strings.Compare("create-token", os.Args[1]) == 0 {
		token, _ := GenerateRandomString(64)
		fmt.Println("Generated token: " + token)
		ioutil.WriteFile("token", []byte(token), 0622)

	} else if len(os.Args) == 2 &&
		strings.Compare("serve", os.Args[1]) == 0 {

		config = getConfiguration()
		db = newDatabase(config.Db, config.DbConnection)
		defer db.Close()

		serve(config.Port)
	}
}

func loadToken() {
	dat, err := ioutil.ReadFile("token")
	if err != nil {
		panic("Couldn't read token file. You can generate a new one with \"SpaceAPI-Server create-token\"\n" + err.Error())
	}
	token = string(dat)
	fmt.Println("Token:\n" + string(dat))
}
