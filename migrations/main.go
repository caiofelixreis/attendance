package main

import (
	"fmt"

	"github.com/caiofelixreis/attendance/pkg/database"
	"github.com/caiofelixreis/attendance/pkg/models"
)

func main() {
	conn, err := database.Postgres.Init()

	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&models.Student{})

	if err != nil {
		panic(err)
	}

	fmt.Println("®️ Migration completed successfully")

	defer database.Postgres.Close()
}
