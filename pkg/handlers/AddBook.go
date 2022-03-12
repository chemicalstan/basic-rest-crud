package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/chemicalstan/go-rest-crud/pkg/mocks"
	"github.com/chemicalstan/go-rest-crud/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	log.Println("E enter")
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the book Mock
	book.Id = rand.Intn(100)

	mocks.Books = append(mocks.Books, book)
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}