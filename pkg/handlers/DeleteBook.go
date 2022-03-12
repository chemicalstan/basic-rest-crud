package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chemicalstan/go-rest-crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for i, book := range mocks.Books {
		if id == book.Id {
			mocks.Books = append(mocks.Books[:i], mocks.Books[i+1:]...)
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(vars["id"] + " deleted successfully")
		break
	}
}
