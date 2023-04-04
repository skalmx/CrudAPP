package main

import (
	"CrudApp/iternal/repository/postgres"
	"CrudApp/iternal/service"
	"CrudApp/iternal/transport/rest"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	 _ "github.com/lib/pq"

)

func main() {
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", "localhost", 5432, "postgres", "postgres", "disable", "123")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Print(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
	}

	defer db.Close()

	lessonsRepository := postgres.NewLessonsRepo(db)
	lessonsService := service.NewLessons(lessonsRepository)
	handler := rest.NewHandler(lessonsService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Init(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// _, err = db.Exec("insert into Products(model, company, price) values ('IphoneXr','apple', 400)")
	// if err != nil{
	// 	fmt.Print(err)
	// } test of db connection and docker containers

}
