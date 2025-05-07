package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/jmgofolio/infrastructure/db"
	"github.com/jmechavez/jmgofolio/infrastructure/logger"
	"github.com/jmechavez/jmgofolio/internal/ports"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Start() {
	router := mux.NewRouter()

	dbPortfolio := getPostgresConnection()
	// Test the connection
	if err := dbPortfolio.Ping(); err != nil {
		logger.Fatal("Failed to ping database", zap.Error(err))
	}
	defer dbPortfolio.Close() // Ensure connection is closed when app terminates

	ph := AppHandler{
		ports.NewPortfolioService(db.NewPortfolioRepository(dbPortfolio)),
	}

	router.HandleFunc("/api/portfolio", ph.PortfolioJSONHandler).Methods("GET")
	router.HandleFunc("/api/htmlportfolio",ph.PorfolioHTTPHandler).Methods("GET")
	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	router.HandleFunc("/contact", ContactHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getPostgresConnection() *sqlx.DB {
	logger.Info("Connecting to PostgreSQL database")

	// Consider using environment variables for these values
	connStr := "host=localhost port=5440 user=admin password=Admin123 dbname=portfolio_db sslmode=disable"

	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL database", zap.Error(err))
	}

	// Configure connection pool settings
	userDb.SetMaxOpenConns(25)
	userDb.SetMaxIdleConns(5)

	logger.Info("Successfully connected to PostgreSQL database")
	return userDb
}
