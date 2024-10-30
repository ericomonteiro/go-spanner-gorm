package main

import (
	"fmt"
	spannergorm "github.com/googleapis/go-gorm-spanner"
	"go-spanner/pkg/entity"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	_ = os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010")

	db, err := gorm.Open(spannergorm.New(spannergorm.Config{
		DriverName: "spanner",
		DSN:        "projects/demo-go-project/instances/simulator-instance/databases/demo-database",
	}), &gorm.Config{PrepareStmt: true})
	if err != nil {
		log.Fatal(err)
	}

	var users []entity.User
	db.Find(&users)

	for _, user := range users {
		fmt.Println(user)
	}
}
