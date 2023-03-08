package main

import (
	"database/sql"

	"github.com/devfullcycle/gointesivo2/internal/infra/database"
	"github.com/devfullcycle/gointesivo2/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() // Executa tudo e depois executa o close

	repository := database.NewOrderRepository(db)

	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}
}
