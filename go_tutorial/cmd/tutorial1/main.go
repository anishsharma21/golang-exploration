package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 5
	var denominator int = 4
	var result, remainder, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Print(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", result, remainder)

	}

	var intarr [3]int32
	fmt.Println()
	intarr[1] = 123
	fmt.Println(intarr[0])
	fmt.Println(intarr[1:3])
	fmt.Println(&intarr[0])
	fmt.Println(&intarr[1])
	fmt.Println(&intarr[2])

	intarr2 := [...]int32{1, 2, 3}
	fmt.Println(intarr2)

	var intslice []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capacity %v", len(intslice), cap(intslice))
	intslice = append(intslice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intslice), cap(intslice))

	var intslice2 []int32 = []int32{8, 9}
	intslice = append(intslice2, intslice...)
	fmt.Println(intslice)

	var intslice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intslice3, len(intslice3), cap(intslice3))
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}