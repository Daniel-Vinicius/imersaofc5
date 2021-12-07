package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Daniel-Vinicius/challenge-golang/adapter/repository"
	"github.com/Daniel-Vinicius/challenge-golang/usecase/create_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepository(db)
	usecase := create_transaction.NewCreateTransactionUseCase(repo)

	input := create_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    100,
	}

	err = usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	var count int
	db.QueryRow(`select COUNT(*) from transactions`).Scan(&count)

	fmt.Println(count, "linhas no banco de dados")
}
