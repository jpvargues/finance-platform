package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/joho/godotenv"
)

type StatusCode struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

var db *sql.DB

/*
connectDB opens a connection pool to Postgres. Note: sql.Open only
validates the connection string and prepares the pool — it doesn't
actually contact the server. Ping() below performs the real handshake
and is what tells us if Postgres is actually reachable.
*/
func connectDB() error {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}
	db, err = sql.Open("pgx", dbURL)
	if err != nil {
		return err
	}
	return db.Ping()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s StatusCode
	s.Status = "ok"

	if err := db.Ping(); err != nil {
		s.Database = "disconnected"
	} else {
		s.Database = "connected"
	}

	userJSON, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(userJSON)
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, relying on real environment variables")
	}

	if err := connectDB(); err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	fmt.Println("Database connected")

	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
