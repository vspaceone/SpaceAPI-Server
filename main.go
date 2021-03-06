package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

var config configuration
var db *gorm.DB

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Wrong input.")
	} else if strings.Compare("create-token", os.Args[1]) == 0 {
		generateToken()
	} else if len(os.Args) == 2 &&
		strings.Compare("serve", os.Args[1]) == 0 {

		config = getConfiguration()
		db = newDatabase(config.Db, config.DbConnection, config.Debug)
		defer db.Close()

		serve(config.Port)
	}
}

func generateToken() {
	token, _ := GenerateRandomString(64)
	fmt.Println("Generated token: " + token)
	ioutil.WriteFile("data/token", []byte(token), 0622)
}
