package main

import (
	"fmt"
	"os"

	"github.com/I1Asyl/task-manager-go/database"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg[1])
	db, err := database.NewConnection(argsWithoutProg[1])
	if err != nil {
		panic(err)
	}
	text := argsWithoutProg[0]
	var dat []byte
	if text == "down" {
		dat, err = os.ReadFile("database_down.pgsql")

	} else if text == "up" {
		dat, err = os.ReadFile("database_up.pgsql")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
	_, err = db.Query(string(dat))
	if err != nil {
		panic(err)
	}

	db.Close()
	os.Exit(0)
}
