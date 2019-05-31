package main
import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"reflect"
	"strconv"
)
type Book struct{
	ID 		int `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Year 	string `json:"year"`
}


var books []Book
func main(){
books = append(books, Book{ID:1,Title:"Golang Pointers",Author:"Mr. Golang",Year:"2010"},
		Book{ID:2,Title:"Goroutines",Author:"Mr. Go. Routine",Year:"2011"},
		Book{ID:3,Title:"Golang routers",Author:"Mr. Router",Year:"2012"},
		Book{ID:4,Title:"Golang concurency",Author:"Mr. Currency",Year:"2013"},
		Book{ID:5,Title:"Golang good oarts",Author:"Mr. Good. Routine",Year:"2014"},
		Book{ID:6,Title:"Beneath Stil Waters",Author:"Mr. Richard",Year:"2015"})



	router := mux.NewRouter()
	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

func getBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request){
	log.Println("GET  book is called")
	params := mux.Vars(r)
	log.Println(params)

	i,_:= strconv.Atoi(params["id"])

	log.Println(reflect.TypeOf(i))

	for _,book := range books {
		if book.ID == i{
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request){
	log.Println("ADD BOOK is called")
}
func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("REMOVE BOOK IS CALLED books is called")
}
func updateBook(w http.ResponseWriter, r *http.Request){
	log.Println("UPDATE BOOK IS CALLED books is called")
}