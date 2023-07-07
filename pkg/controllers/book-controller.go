package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MarselBisengaliev/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	BookModel := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(BookModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b := BookModel.CreateBook()
	res, _:= json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	err := json.NewDecoder(r.Body).Decode(updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, db := models.GetBookById(bookId)
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}