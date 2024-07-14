package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Book struct {
	title string
	author string
	pages int
}

func bookDetails(book Book) {
	fmt.Printf("Book name: %s\nAuthor: %s\nPages: %v", book.title, book.author, book.pages)
	fmt.Println()
}

type Student struct {
	name string
	age int
	grades []int
}

func (student Student) AverageGrade() float64 {
	var sum int = 0
	for _, grade := range student.grades {
		sum += grade
	}
	average := float64(sum) / float64(len(student.grades))
	return average
} 

func (student Student) isPassing() bool {
	return student.AverageGrade() > 60
}

func greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func calculateBMI(height, weight float64) float64 {
	return float64(weight / (height * weight))
}

func main() {
	p := Person{Name: "anish sharma", Age: 25}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
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

	for i := 1; i < 11; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Println(greet(name))
	fmt.Println(calculateBMI(height, 67))

	book := Book{title: "harry potter", author: "JK", pages: 300}

	bookDetails(book)

	student := Student{name: "anish", age: 27, grades: []int{45, 65, 78, 89, 23}}
	averageGrade := student.AverageGrade()
	isPassing := student.isPassing()
	fmt.Printf("%v, %v", averageGrade, isPassing)
	fmt.Println()
}