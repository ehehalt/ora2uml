package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ehehalt/ora2uml"
	_ "github.com/godror/godror"
)

func main() {
	fmt.Println("ora2uml starter")

	if len(os.Args) < 2 {
		fmt.Println("ora2uml configfile.json")
		os.Exit(0)
	}
	configFileName := os.Args[1]

	fmt.Println(configFileName)

	config, err := ora2uml.Read(configFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("UserId =", config.User.UserId, "Password =", config.User.Password)
	fmt.Println("Host =", config.Database.Host, "Port =", config.Database.Port, "ServiceName =", config.Database.ServiceName)
	for _, value := range config.Tables {
		fmt.Println("Table =", value.Name)
	}

	db, err := sql.Open("godror", config.ConnectionString())
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Connection successful")
	db.Close()
}
