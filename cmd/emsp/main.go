package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)
}

func main() {
	err := StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
