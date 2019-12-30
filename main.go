package main

import (
	"expense-tracker/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	registerRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", auth.Login).Methods(http.MethodGet)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
