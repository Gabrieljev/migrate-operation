package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mongodb"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/database/postgres"

	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	fmt.Println("testing migration db with mongodb, postgres, and mysql")
	m, err := migrate.New("file://./filesmongo", "mongodb://@localhost:27017/admin/testing")
	if err != nil {
		log.Fatal("mongodb has been failed to connect", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal("mongodb has been failed migration", err)
	}
	m.Close()

	m, err = migrate.New("file://./filespostgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("postgres has been failed to connect", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal("postgres has been failed migration", err)
	}
	m.Close()

	m, err = migrate.New("file://./filesmysql", "mysql://root:test@tcp(localhost:3306)/sys?multiStatements=true")
	if err != nil {
		log.Fatal("mysql has been failed to connect", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal("mysql has been failed migration", err)
	}
	m.Close()
	fmt.Println("succes migrating")

}
