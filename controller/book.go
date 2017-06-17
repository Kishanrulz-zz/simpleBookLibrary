package controllers

import (
	"ExamplesDemo/Lib/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//CreateBook API creates a book
func (api Api) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	book := &models.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Bad request")
		return
	}

	uj, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Bad request")
		return
	}
	//calls the create method of books
	if err := book.Create(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", uj)
		return
	} else {
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", uj)
		return
	}
	fmt.Fprintf(w, "%s\n", "Book created successfully")
}

//GetBook fetches a book based on the ID
func (api Api) GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := &models.Book{
		ID: bson.ObjectIdHex(p.ByName("id")),
	}

	if err := book.Get(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", book)
		return
	}

	uj, err := json.Marshal(book)
	if err != nil {
		fmt.Fprintf(w, "%s\n", book)
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%s\n", uj)
}

//UpdateBook updates a book
func (api Api) UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := &models.Book{}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", book)
		return
	}

	book.ID = bson.ObjectIdHex(p.ByName("id"))

	if err := book.Update(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", book)
		return
	}

	uj, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Invalid input")
		return
	}
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteBook deletes a book
func (api Api) DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := &models.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Invalid input")
		return
	}
	if bson.IsObjectIdHex(p.ByName("id")) == false {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Invalid bookID")
		return
	}
	book.ID = bson.ObjectIdHex(p.ByName("id"))

	if err := book.Delete(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	uj, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	fmt.Fprintf(w, "%s\n", uj)
}

//IssueBook issues a book to a user
func (api Api) IssueBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	issue := &models.IssueBook{}

	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Sorry, invalid input")
		return
	}

	book := &models.Book{
		ID: bson.ObjectIdHex(p.ByName("id")),
	}

	if err := book.Get(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	if bson.IsObjectIdHex(issue.UserID) == false {
		fmt.Fprintf(w, "%s\n", "Sorry, invalid userID")
		return
	}

	//verifying whether a user exists or not
	if err := isValid(bson.ObjectIdHex(issue.UserID), api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err.Error() == "Invalid user" {
			fmt.Fprintf(w, "%s\n", "Sorry, user already deleted")
			return
		} else {
			fmt.Fprintf(w, "%s\n", "Sorry, user not found")
			return
		}
	}
	//updating the book meta with the userID to whom the book has been issued
	book.UserID = append(book.UserID, bson.ObjectIdHex(issue.UserID))
	if err := book.Update(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	fmt.Fprintf(w, "%s\n", "Book issued successfully")
}
