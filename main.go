package main

import (
	"log"

	"github.com/caiofelixreis/attendance/pkg/database"
	"github.com/caiofelixreis/attendance/pkg/models"
	repository "github.com/caiofelixreis/attendance/pkg/repositories"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: No .env file found")
	}

	database.Postgres.Init()
	defer database.Postgres.Close()

	caio := models.Student{
		ID: "3",
	}

	conn := database.Postgres.GetConn()

	studentRepository := repository.NewStudentRepository(conn)

	err := studentRepository.Create(&caio)

	if err != nil {
		panic(err)
	}
	database.Postgres.PrintSchema()
}
