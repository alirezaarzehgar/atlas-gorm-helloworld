package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
}

func main() {
	pg, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=crm port=5432 sslmode=disable TimeZone=Asia/Tehran"))
	if err != nil {
		log.Fatal(err)
	}

	pg.
}
