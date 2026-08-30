package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/joho/godotenv"
)

type StatusCode struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

type ETF struct {
	ID             int    `json:"id"`
	Ticker         string `json:"ticker"`
	Name           string `json:"name"`
	ISIN           string `json:"isin"`
	IsAccumulating bool   `json:"is_accumulating"`
	Category       string `json:"category"`
}

var db *sql.DB

func respondError(w http.ResponseWriter, statusCode int, logMsg string, err error) {
	log.Println(logMsg, err)
	http.Error(w, logMsg, statusCode)
	return
}

func getAllETFs() ([]ETF, error) {
	rows, err := db.Query("SELECT id, ticker, name, isin, is_accumulating, category FROM etfs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	etfs := []ETF{}
	for rows.Next() {
		var e ETF
		err := rows.Scan(&e.ID, &e.Ticker, &e.Name, &e.ISIN, &e.IsAccumulating, &e.Category)
		if err != nil {
			return nil, err
		}
		etfs = append(etfs, e)
	}

	return etfs, nil
}

// connectDB opens a connection pool to Postgres. Note: sql.Open only
// validates the connection string and prepares the pool — it doesn't
// actually contact the server. Ping() below performs the real handshake
// and is what tells us if Postgres is actually reachable.
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

	statusJSON, err := json.Marshal(s)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "JSON encoding error:", err)
		return
	}
	w.Write(statusJSON)
}

func etfs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	etfs, err := getAllETFs()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "database query error", err)
		return
	}

	etfJSON, err := json.Marshal(etfs)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "JSON encoding error:", err)
		return
	}

	w.Write(etfJSON)
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
	http.HandleFunc("/etfs", etfs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
