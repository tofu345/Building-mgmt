package main

import (
	"log"
	"net/http"
	"os"
	"time"

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
	r.Use(internal.LoggingMiddleware)
	r.Use(allowCorsMiddleware)

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

func allowCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range internal.ALLOWED_HOSTS {
			w.Header().Set("Access-Control-Allow-Origin", v)
		}
		next.ServeHTTP(w, r)
	})
}
