package main

import "fmt"

type Person struct {
	Name string
	Age int
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
}