package main

import (
	"expense-tracker/auth"
	"expense-tracker/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload" // load in environmental variables from .env
)

func main() {
	db := database.InitializeDb()
	defer db.Close()
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	registerRoutes(db, router)
	modifiedRouter := handlers.LoggingHandler(os.Stdout, router)
	modifiedRouter = handlers.CompressHandler(modifiedRouter)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", modifiedRouter))
}

func registerRoutes(db *gorm.DB, router *mux.Router) {
	router.HandleFunc("/", auth.Login).Methods(http.MethodGet)
	router.HandleFunc("/signup", auth.Signup).Methods(http.MethodPost)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
