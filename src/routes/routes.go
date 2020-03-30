package routes

import (
	"fmt"
	"log"
	"net/http"

	"../articles"
	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)

// RegisterRoutes regiters routes main entrypoint
func RegisterRoutes() {
	log.Printf("Registering routes")
	myRouter.Use(commonMiddleware)
	fetchAllRoutes()

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving health chceck")
	fmt.Fprintf(w, "Server running correctly")
}

func fetchAllRoutes() {
	myRouter.HandleFunc("/health", healthCheck)
	myRouter.HandleFunc("/articles", articles.AllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", articles.ArticleByID)
	myRouter.HandleFunc("/articles", articles.CreateArticle).Methods("POST")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("fire 1")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
