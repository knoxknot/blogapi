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

// program entry point
func main() {

}