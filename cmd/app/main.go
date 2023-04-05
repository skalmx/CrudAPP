package main

import (
	"CrudApp/iternal/repository/postgres"
	"CrudApp/iternal/service"
	"CrudApp/iternal/transport/rest"
	// "context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "os/signal"
	"time"
	// "syscall"
	_ "github.com/lib/pq"
)

func main() {
	// connection to db
	// to do create config for connections to server and db
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", "localhost", 5432, "postgres", "postgres", "disable", "123")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Print(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
	}

	defer db.Close()
	// init dependencies
	lessonsRepository := postgres.NewLessonsRepo(db)
	lessonsService := service.NewLessons(lessonsRepository)
	handler := rest.NewHandler(lessonsService)
	//start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Init(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	// go func (){
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	// }()
	
	// quit := make(chan os.Signal, 1) 
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// <-quit

	// const timeout = 5 * time.Second

	// ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	// defer shutdown()
	
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Print("failed to stop server:")
	// }
}
