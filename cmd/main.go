package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/StevenD2002/Medical-Bill-Insights/internal/db/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error converting port to int")
	}

	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("DBPASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// use the new repository
	procedureRepo := repository.NewProcedureRepository(db)

	// use the repository to get a procedure by code
	procedure, err := procedureRepo.GetProcedurePricesByCode("99203")

	if err == nil {
		fmt.Println(procedure.Code)
	} else {
		fmt.Println(err)
	}
	fmt.Println("Program exited successfully")
}
