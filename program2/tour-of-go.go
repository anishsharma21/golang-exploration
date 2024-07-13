package main

import "fmt"

func main() {
	defer fmt.Println("done")
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	newsum := 1
	for ; newsum < 1000; {
		newsum += newsum
	}
	fmt.Println(newsum)
}