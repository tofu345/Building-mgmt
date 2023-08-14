package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tofu345/Building-mgmt-backend/internal"
	"github.com/tofu345/Building-mgmt-backend/scripts"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "shell":
			scripts.Shell()
		default:
			log.Fatalf("Unknown verb: %v", os.Args[1])
		}
		return
	}

	port := "127.0.0.1:8000"
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	internal.RegisterRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listening on ", port)
	log.Fatal(srv.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
