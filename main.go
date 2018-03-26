package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"

)
//Book Struct
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
} 
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

} 
// init books struct
var books []Book 
//get all books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)	
}
// get single book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
for _, item := range books {
	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
}
json.NewEncoder(w).Encode(&Book{})
	
}
//create book
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
 book.ID = strconv.Itoa(rand.Intn(1000000))
 books = append(books, book)
 json.NewEncoder(w).Encode(book)
	
}
//update book
func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
 book.ID = strconv.Itoa(rand.Intn(1000000))
 books = append(books, book)
 json.NewEncoder(w).Encode(book)
  return 
		}
	}
	json.NewEncoder(w).Encode(books)

	
}
//delete book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(books)

}
func main(){
	// initialize router
	r := mux.NewRouter()
	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn:"44412", Title: "book one", Author: &Author{
		Firstname: "brian", Lastname: "mituka"}})
	books = append(books, Book{ID: "2", Isbn:"34521", Title: "book two", Author: &Author{
		Firstname: "chimie", Lastname: "mots"}})
	books = append(books, Book{ID: "3", Isbn:"456789", Title: "book three", Author: &Author{
		Firstname: "Alex", Lastname: "tito"}})
	
	

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
	
}