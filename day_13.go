package main

import "fmt"

type Book struct {
	title  string
	author string
}

type MyBook struct {
	Book
	price int
}

func NewMyBook(title, author string, price int) MyBook {
	book := Book{title, author}
	return MyBook{book, price}
}

func (book MyBook) Display() {
	fmt.Println("Title:", book.title)
	fmt.Println("Author:", book.author)
	fmt.Println("Price:", book.price)
}

func main() {
	newNovel := NewMyBook("The Alchemist", "Paulo Coelho", 248)
	newNovel.Display()
}
