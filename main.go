package main

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

// program entry point
func main() {

}