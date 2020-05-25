package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func changeBook(books *Books) {
	books.title = "book1_change"
}

func (r *Books) getName() string {
	r.title = "book2_change"
	return r.title
}

func main() {
	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.bookID = 1
	changeBook(&book1)
	fmt.Println(book1)
	fmt.Println(book1.getName())
	fmt.Println(book1)

	fmt.Println(Books{"a", "b", "c", 11})
	fmt.Println(Books{title: "a1", author: "b1", subject: "c1", bookID: 21})
	fmt.Println(Books{author: "c1", bookID: 3333})
}
