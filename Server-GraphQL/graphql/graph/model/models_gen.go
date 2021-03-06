// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	Author *User  `json:"author"`
}

type NewBook struct {
	Title  string `json:"title"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
