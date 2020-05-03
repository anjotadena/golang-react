package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Infof("uri: %s", request.RequestURI)

	response.Write([]byte("Hello, World!"))
}

func main() {
	fmt.Println("Running Go API Web services...")

	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedHeaders([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedHeaders([]string{"*"})

	// ROUTES
	router.HandleFunc("/", RootEndpoint).Methods("GET")

	http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(router))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World AGAIN")
// }

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }
