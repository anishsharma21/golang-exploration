package main

import "fmt"

func main() {
	var age int = 17
	var height float64 = 175
	var isStudent bool = true
	var name string = "anish"
	fmt.Println(age, height, isStudent, name)

	if age < 18 {
		fmt.Println("Underage")
	} else {
		fmt.Println("Adult")
	}

	switch name {
	case "anish":
		fmt.Println("Hello Anish!")
	default:
		fmt.Println("Hello guest!")
	}
}