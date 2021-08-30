package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string `json:"-"`
}

type Books struct {
	BookID     int
	BookName   string
	BookAuthor string
}

type SliceUsers []User

type SliceBooks []Books

var VarSliceUsers SliceUsers
var VarSliceBooks SliceBooks

func init() {

	VarSliceUsers = []User{
		User{
			ID:       1,
			Username: "user1",
			Password: "pass1",
		},
		User{
			ID:       2,
			Username: "user2",
			Password: "pass2",
		},
		User{
			ID:       3,
			Username: "user3",
			Password: "pass3",
		},
	}

	VarSliceBooks = []Books{
		Books{
			BookID:     1,
			BookName:   "book1",
			BookAuthor: "auther1",
		},
		Books{
			BookID:     2,
			BookName:   "book2",
			BookAuthor: "auther2",
		},
		Books{
			BookID:     3,
			BookName:   "book3",
			BookAuthor: "auther3",
		},
	}

}

func main() {

	http.HandleFunc("/books", listBooks)
	http.HandleFunc("/users", listUsers)

	LASerr := http.ListenAndServe(":9091", nil)

	if LASerr != nil {
		log.Fatal(LASerr)
	}
}

func listBooks(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(VarSliceBooks)

	if err != nil {
		log.Panic(err)
	}

}

func listUsers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(VarSliceUsers)

	if err != nil {
		log.Panic(err)
	}

}
