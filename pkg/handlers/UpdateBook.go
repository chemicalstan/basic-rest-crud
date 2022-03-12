package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/chemicalstan/go-rest-crud/pkg/mocks"
	"github.com/chemicalstan/go-rest-crud/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	for i, book := range mocks.Books {
		if id == book.Id {
			book.Author = updatedBook.Author
			book.Title = updatedBook.Title
			book.Desc = updatedBook.Desc
			mocks.Books[i] = book

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Book updated")
			break
		}
	}
}
