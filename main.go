package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/joaosp7/GoClassicTodo/internal/handlers"
	"github.com/joaosp7/GoClassicTodo/internal/repository"
	"github.com/joaosp7/GoClassicTodo/internal/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading dotEnv.")
	}
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "banana"),
		getEnv("DB_NAME", "todo"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	log.Println("Successfully connected to database!")

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	handler := handler.NewHandler(userService)

	http.HandleFunc("/user", handler.CreateUser)
	http.HandleFunc("/hello", hello)
	fmt.Println("Server up and running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
