package main

import (
	"fmt"
)

type Book struct {
	authorName    string
	pageNumber    int
	chapter       int
	numCopies     int
	price         float64
	pubYear       int
	isbn          int
	goodCondition bool
}

type Author struct {
	fullName        string
	notableSeries   string
	age             int
	totalSold       int
	numPublications int
}

func NewAuthor(fName string, notable string, aAge int, sold int, num int) *Author {
	temp := Author{fName, notable, aAge, sold, num}
	return &temp
}

func changeName(name string, author *Author) *Author {
	author.fullName = name
	return author
}
func printAuthor(author *Author) {
	fmt.Println("Full Name: ", author.fullName)
	fmt.Println("Notable Series: ", author.notableSeries)
	fmt.Println("Age: ", author.age)
	fmt.Println("Total Copies Sold: ", author.totalSold)
	fmt.Println("Number of Series Published: ", author.numPublications)
}

func returnHighestAge(book []*Book) *Book {
	var largest = book[0]
	for i := 1; i < len(book); i++ {
		if book[i].pubYear > largest.pubYear {
			largest = book[i]
		}
	}
	return largest
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
	println("Book author: ", book.authorName)
	println("Pages: ", book.pageNumber)
	println("Chapters: ", book.chapter)
	println("Number of copies: ", book.numCopies)
	priceStr := fmt.Sprintf("%.2f", book.price)
	println("Price: ", priceStr)
	println("Year of publication: ", book.pubYear)
	println("ISBN: ", book.isbn)
}

func manage_sell_price(book *Book) {
	if book.goodCondition {
		increasePrice(book, 2)
	} else if !book.goodCondition {
		decreasePrice(book, 5)
	} else {
		fmt.Println("Invalid condition")
	}
}

func main() {
	JKRowling := NewAuthor("JKRowling", "Harry Potter", 52, 100000, 8)
	printAuthor(JKRowling)
	harryPotter := NewBook("J.K. Rowling", 500, 20, 1000000, 19.99, 1997, 9788700631625, true)
	println("before price change")
	printBookInfo(harryPotter)
	println("After price change")
	println()
	increasePrice(harryPotter, 12.99)
	printBookInfo(harryPotter)
	fmt.Println("End of main method")
}
