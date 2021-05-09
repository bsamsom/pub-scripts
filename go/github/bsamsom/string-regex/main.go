package main

import (
	"fmt"
	"regexp"
)

// function, input of type string, return type string.
func regex(input string) string {
	email := regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	shortPhone := regexp.MustCompile(`^[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)
	longPhone := regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)
	switch {
	case email.MatchString(input):
		return "email"
	case shortPhone.MatchString(input):
		return "shortPhone"
	case longPhone.MatchString(input):
		return "longPhone"
	default:
		return ""
	}
}

func main() {
	// var declaration type 1
	/*
		var name string = "ben"
		var age int32 = 32
		fmt.Println("Name", name, "Age", age)
	*/
	// var declaration type 2
	/*
		name1 := "ben1"
		age1 := 25
		fmt.Println("Name1", name1, "Age1", age1)

		// find varible types
		fmt.Printf("%s is a %T\n", name, name)
		fmt.Printf("%d is a %T\n", age, age)
	*/
	arrayExample := [...]string{"test@test.ca", "123-4567", "123-456-7890"}
	fmt.Println("arrayExample has a length of:", len(arrayExample))

	for index, value := range arrayExample {
		fmt.Println(index, regex(value))
	}

}
