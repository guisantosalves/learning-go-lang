package main

import "fmt"

type User struct {
	FirstName, LastName string
	age                 int
}

func main() {
	// map is the same of dicts in py

	// key: value -> string: int / key value pair data
	grades := make(map[string]int)

	grades["Harry potter"] = 23
	grades["john doe"] = 5
	grades["guizaodozap"] = 20

	// fmt.Println(grades)

	// delete(grades, "john doe")

	// fmt.Println(grades)

	useDataBase := make(map[string]*User)

	// creating the objects of the structs
	userOne := User{
		FirstName: "guigui",
		LastName:  "dozap",
		age:       20,
	}
	userTwo := User{
		FirstName: "beaa",
		LastName:  "dozapi",
		age:       10,
	}
	userThree := User{FirstName: "euzin",
		LastName: "dozapp",
		age:      96,
	}

	// pass the references of objects
	useDataBase["1"] = &userOne
	useDataBase["2"] = &userTwo
	useDataBase["3"] = &userThree

	fmt.Println(useDataBase["1"].FirstName)
	fmt.Println(useDataBase["2"].FirstName)
	fmt.Println(useDataBase["3"].FirstName)

	// for in maps works in a different way the args are key and value
	for key, value := range useDataBase {
		fmt.Println("user id: ", key, " user: ", value.FirstName)
	}
}
