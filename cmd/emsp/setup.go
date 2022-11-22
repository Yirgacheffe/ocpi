package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	sql "github.com/Yirgacheffe/ocpi/db"
	"github.com/Yirgacheffe/ocpi/versions"
	"github.com/gorilla/mux"

	"github.com/Yirgacheffe/ocpi/middleware"
)

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	} else {
		return fallback
	}
}

func StartServer() error {

	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPazz := getEnv("DB_PAZZ", "postgres")
	dbName := getEnv("DB_NAME", "ocpi_emsp")

	db, err := sql.ConnectSQL(dbHost, dbPort, dbUser, dbPazz, dbName)
	if err != nil {
		return err
	}

	addr := getEnv("ADDR", "localhost:9090")

	// version endpoint
	vRepo := versions.NewVersionRepo(db.DB)
	v := versions.NewHandler(vRepo)

	router := mux.NewRouter()
	api := router.PathPrefix("/ocpi/emsp").Subrouter()
	api.HandleFunc("/versions", v.ListVersions).Methods("GET", "OPTIONS")

	// middleware
	router.Use(middleware.WithMetrics)
	router.Use(mux.CORSMethodMiddleware(router))

	log.Printf("server is running at: %s", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		return err
	}

	return nil // should start with no error happened

}
