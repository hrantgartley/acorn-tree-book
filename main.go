package main

import (
	"fmt"
)

type Book struct {
	author        string
	pageNumber    int
	chapter       int
	numCopies     int
	price         float64
	pubYear       int
	isbn          int
	goodCondition bool
}

func NewBook(author string, pageNumber int, chapter int, numCopies int, price float64, pubYear int, isbn int, cond bool) *Book {
	temp := Book{author, pageNumber, chapter, numCopies, price, pubYear, isbn, cond}
	return &temp
}

func addBookToArray(book *Book, bookArray []*Book) []*Book {
	bookArray = append(bookArray, book)
	return bookArray
}

func removeBookFromArray(book *Book, bookArray []*Book) []*Book {
	for i, b := range bookArray {
		if b == book {
			bookArray = append(bookArray[:i], bookArray[i+1:]...)
			break
		}
	}
	return bookArray
}

func increasePrice(book *Book, price float64) {
	book.price += price
}

func decreasePrice(book *Book, price float64) {
	book.price -= price
}

func isbnValid(book *Book) bool {
	return book.isbn > 1000000000
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
	harryPotter := NewBook("J.K. Rowling", 500, 20, 1000000, 19.99, 1997, 9788700631625, true)
	println("before price change")
	printBookInfo(harryPotter)
	println("After price change")
	println()
	increasePrice(harryPotter, 12.99)
	printBookInfo(harryPotter)
}
