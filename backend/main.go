package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tofu345/Building-mgmt-backend/scripts"
	"github.com/tofu345/Building-mgmt-backend/src/middleware"
	"github.com/tofu345/Building-mgmt-backend/src/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "shell":
			if len(os.Args) >= 2 {
				scripts.Shell(os.Args[2:]...)
			} else {
				scripts.Shell()
			}
		default:
			log.Fatalf("Unknown verb: %v", os.Args[1])
		}
		return
	}

	port := "127.0.0.1:8000"
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.AllowCors)

	routes.RegisterRoutes(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins(strings.Split(os.Getenv("ALLOWED_HOSTS"), ","))
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      handlers.CORS(headersOk, originsOk, methodsOk)(r),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listening on ", port)
	log.Fatal(srv.ListenAndServe())
}
