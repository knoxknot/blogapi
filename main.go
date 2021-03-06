package main

// packages that enables our task
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// define the data type
type article struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Content string `json:"Content"`
}

// A slice mock data of all articles described by the article struct type
var articles = []article{
	article{ID: "1",Title: "Introduction to Go", Content: "This is a brief about Go and its awesomeness"},
	article{ID: "2",Title: "The simplicity of Go", Content: "There is so much Awesomeness in Go"},
	article{ID: "3",Title: "Go Mastery", Content: "If you love simplicity then you should master Go"},
}

// implement a retrival of all existing articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
	w.WriteHeader(http.StatusOK)
}

// implement a single article retrival by id
func getAnArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, article := range articles {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return	
		}
	}
	json.NewEncoder(w).Encode(&article{})
	w.WriteHeader(http.StatusOK)
}

func createAnArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newArticle article
	_ = json.NewDecoder(r.Body).Decode(&newArticle)
	articles = append(articles, newArticle) //{"ID": "4","Title": "Go: Clean, Readable and Fast","Content": "It is clean in design, easy to read  and outputs very fast."}
	json.NewEncoder(w).Encode(newArticle)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"success": "article created"}`))
}

// implement the updating of a given article by id
func updateAnArticle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var updatedArticle article
	_ = json.NewDecoder(r.Body).Decode(&updatedArticle)
	params := mux.Vars(r)
	for index, article := range articles {
		if article.ID == params["id"] {  
			updatedArticle.ID = params["id"] 
			article.Title = updatedArticle.Title
			article.Content = updatedArticle.Content
			articles = append(articles[:index], updatedArticle) //{"ID":"4","Title":"Go is Clean and Readable","Content":"It is clean in design, easy to read  and outputs very fast."}
			json.NewEncoder(w).Encode(updatedArticle)
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(`{"success": "article updated"}`))
		}
	}
}

// implements the deletion of a specific article by id
func deleteAnArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, article := range articles {
		if article.ID == params["id"] {
			articles = append(articles[:index],articles[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(&article{})
	w.WriteHeader(http.StatusOK)
}

// program entry point
func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/articles", getAllArticles).Methods("GET")
	api.HandleFunc("/articles/{id}", getAnArticle).Methods("GET")
	api.HandleFunc("/articles", createAnArticle).Methods("POST")
	api.HandleFunc("/articles/{id}", updateAnArticle).Methods("PUT")
	api.HandleFunc("/articles/{id}", deleteAnArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",r))
}