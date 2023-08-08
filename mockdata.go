package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	AuthorID int    `json:"authorId"`
}

var authorsData = []Author{
	{1, "J. K. Rowling"},
	{2, "J. R. R. Tolkien"},
	{3, "Brent Weeks"},
}

var booksData = []Book{
	{1, "Harry Potter and the Chamber of Secrets", 1},
	{2, "Harry Potter and the Prisoner of Azkaban", 1},
	{3, "Harry Potter and the Goblet of Fire", 1},
	{4, "The Fellowship of the Ring", 2},
	{5, "The Two Towers", 2},
	{6, "The Return of the King", 2},
	{7, "The Way of Shadows", 3},
	{8, "Beyond the Shadows", 3},
}
