/* Question 3
In Go, write a function named topAuthors that accepts a slice of Book structs and returns a
map[string]int where the keys are the author names and the values are the sum of the sales of
all their books. The function should only include authors in the result map if the total sales of all
their books are at least 10000.
The Book struct is defined as follows:

type Book struct {
	Title string
	Author string
	Sales int
} */

//go:build practice3
//run with $ go run -tags practice3 practice3.go
package main

import (
	"fmt"
)

type Book struct {
	Title string
	Author string
	Sales int
}

func topAuthors(books []Book) map[string]int {
	authorSales := make(map[string]int)
	for _, book := range books {
		authorSales[book.Author] += book.Sales
	}
	for author, sales := range authorSales {
		if sales < 10000 {
			delete(authorSales, author)
		}
	}
	return authorSales
}

func main() {
	books1 := []Book{
		Book{
			Title: "Book 1",
			Author: "Author 1",
			Sales: 5000,
		},
		Book{
			Title: "Book 2",
			Author: "Author 2",
			Sales: 10000,
		},
		Book{
			Title: "Book 3",
			Author: "Author 1",
			Sales: 6000,
		},
		Book{
			Title: "Book 4",
			Author: "Author 3",
			Sales: 3000,
		},
		Book{
			Title: "Book 5",
			Author: "Author 1",
			Sales: 2000,
		},
	}

	books2 := []Book{
		Book{
			Title: "The Great Escape",
			Author: "John Doe",
			Sales: 4500,
		},
		Book{
			Title: "The Lost World",
			Author: "Michael Crichton",
			Sales: 20000,
		},
		Book{
			Title: "The Final Countdown",
			Author: "John Doe",
			Sales: 8000,
		},
		Book{
			Title: "The Tale of Two Cities",
			Author: "Charles Dickens",
			Sales: 5000,
		},
		Book{
			Title: "Under the Dome",
			Author: "Stephen King",
			Sales: 25000,
		},
		Book{
			Title: "Pet Sematary",
			Author: "Stephen King",
			Sales: 8000,
		},
	}
	fmt.Println(topAuthors(books1)) // map[string]int{ "Author 1": 13000, "Author 2": 10000, }
	fmt.Println(topAuthors(books2)) // map[string]int{ "John Doe": 12500, "Michael Crichton": 20000, "Stephen King": 33000, }
}