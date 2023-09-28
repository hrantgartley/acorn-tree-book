package main

import (
	"fmt"
)

type Book struct {
	author     string
	pageNumber int
	chapter    int
	numCopies  int
	price      float64
	pubYear    int
	isbn       int
}

func NewBook(author string, pageNumber int, chapter int, numCopies int, price float64, pubYear int, isbn int) *Book {
	temp := Book{author, pageNumber, chapter, numCopies, price, pubYear, isbn}
	return &temp
}

func addBookToArray(book *Book, bookArray []*Book) []*Book {
	bookArray = append(bookArray, book)
	return bookArray
}

func printBookInfo(book *Book) {
	println("Book author: ", book.author)
	println("Pages: ", book.pageNumber)
	println("Chapters: ", book.chapter)
	println("Number of copies: ", book.numCopies)
	priceStr := fmt.Sprintf("%.2f", book.price)
	println("Price: ", priceStr)
	println("Year of publication: ", book.pubYear)
	println("ISBN: ", book.isbn)
}

func main() {
	harryPotter := NewBook("J.K. Rowling", 500, 20, 1000000, 19.99, 1997, 9788700631625)
	printBookInfo(harryPotter)
}
