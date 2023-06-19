package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer *Writer `json:"writer"`
}

type Writer struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	params := mux.Vars(r)
	for index, value := range books {
		if value.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, value := range books {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, value := range books {
		if value.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Haya aur Pakdamini", Writer: &Writer{Firstname: "Hazrat Maulana Hafiz Peer Zulfiqar", Lastname: "Naqsbandi Mujadidi D.B."}})
	books = append(books, Book{ID: "2", Title: "Ba Adab Ba Naseeb", Writer: &Writer{Firstname: "Hazrat Maulana Hafiz Peer Zulfiqar", Lastname: "Naqsbandi Mujadidi D.B."}})
	books = append(books, Book{ID: "3", Title: "Allah Se Maango", Writer: &Writer{Firstname: "Hazrat Maulana Syyed Mohammed Talha", Lastname: "Qasmi Naqsbandi Mujadidi D.B."}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Starting server at port 3000:")
	log.Fatal(http.ListenAndServe(":3000", r))
}
