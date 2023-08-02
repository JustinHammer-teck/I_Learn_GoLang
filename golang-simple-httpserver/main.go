package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloResponse)
	http.HandleFunc("/book", getBooks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloResponse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.Write([]byte("<h1 style='color: red;'>Hello From the Server Side</h1>"))
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	book := Book{
		Title:  "The power of habit",
		Author: "",
		Price:  5,
	}

	err := json.NewEncoder(writer).Encode(book)
	if err != nil {
		return
	}
}
