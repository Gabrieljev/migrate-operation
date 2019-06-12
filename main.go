package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mongodb"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	fmt.Println("test")
	m, err := migrate.New("file://./files", "mongodb://@localhost:27017/admin/testing")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

}
