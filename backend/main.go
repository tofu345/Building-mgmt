package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	routes.RegisterRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("ALLOWED_HOSTS"), ","),
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	srv := &http.Server{
		Handler:      c.Handler(r),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listening on ", port)
	log.Fatal(srv.ListenAndServe())
}
