package articles

import (
	"fmt"
	"math/rand"
	"net/http"

	"../utils"
	mux "github.com/gorilla/mux"

	"encoding/json"
	"log"
	"strconv"
)

type articles struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var articlesList []articles = []articles{
	{ID: 1, Name: "First", Description: "Random"},
	{ID: 2, Name: "Second", Description: "Random"},
	{ID: 3, Name: "Third", Description: "Random"},
	{ID: 4, Name: "Fourth", Description: "Random"},
}
var random string = "Yoo"

// AllArticles fetches all the articles
func AllArticles(w http.ResponseWriter, r *http.Request) {
	log.Printf("Fetching all articles")
	json.NewEncoder(w).Encode(articlesList)
}

// ArticleByID fetch one article based on id
func ArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("error ", err)
	}
	log.Printf("Fetching one article ID: " + strconv.Itoa(requestID))
	for _, article := range articlesList {
		if article.ID == requestID {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(utils.Errors{ErrorID: "404", ErrorMsg: "Article for id " + strconv.Itoa(requestID) + " not found"})
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article articles
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = rand.Intn(100)
	fmt.Println(article)
	articlesList = append(articlesList, article)
	json.NewEncoder(w).Encode(articlesList)
}
