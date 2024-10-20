package basic_tutorial

import "fmt"

type Book struct {
	name   string
	year   int
	author string
}

func Library() {
	newBook := Book{name: "Romper el circulo", year: 2023, author: "No sabe no responde"}

	fmt.Println(newBook.name)
	fmt.Println(newBook.year)
	fmt.Println(newBook.author)
}
