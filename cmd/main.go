package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/StevenD2002/Medical-Bill-Insights/internal/db/repository"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "prisma"
	dbname   = "prisma"
	password = "prisma"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// use the new repository
	procedureRepo := repository.NewProcedureRepository(db)

	// use the repository to get a procedure by code
	procedure, err := procedureRepo.GetProcedureByCode("19120")

	if err == nil {
		fmt.Println(procedure.Code)
	} else {
		fmt.Println(err)
	}
	fmt.Println("Program exited successfully")
}
