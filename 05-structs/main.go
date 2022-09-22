package main

import "fmt"

// creating struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Country   string
}

type Book struct {
	Title  string
	Author string
}

func main() {
	personOne := Person{
		FirstName: "Frodo",
		LastName:  "Baggins",
		Country:   "Shire",
	}

	bookOne := Book{
		Author: "guizaodozap",
		Title:  "topper",
	}

	fmt.Println("PersonOne: ", personOne.getFullName())

	fmt.Println("BookOne tittle: ", bookOne.Title)

	//pass for reference
	bookOne.changeTitle("harry pooota")
	fmt.Println("BookOne tittle: ", bookOne.Title)
}

// getting Book by reference not value
// func changeTitle(book *Book, newTitle string) {
// 	book.Title = newTitle
// }

// creating a method to book
func (b *Book) changeTitle(newTitle string) {
	b.Title = newTitle
}

// the return is after the name
// this is a method from Person struct
func (p *Person) getFullName() string {
	return p.FirstName + " " + p.LastName
}
