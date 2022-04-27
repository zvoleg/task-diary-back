package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/zvoleg/task-diary-back/internal"
)

func main() {
	connString := fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		5432,
		"task-diary-dev",
		"postgres",
		"test",
	)
	fmt.Println(connString)

	db, err := sqlx.Open("pgx", connString)
	if err != nil {
		log.Fatal(err)
	}

	app := internal.NewApplication(db)
	app.Init()
	app.Run()
}
