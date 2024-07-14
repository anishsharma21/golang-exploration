package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
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

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":34, "Sarah":32}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["John"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Name doesn't exist in the map")
	}
	fmt.Println()
	for name, age := range myMap2 {
		fmt.Printf("Name: %s\nAge: %v\n", name, age)
	}
	fmt.Println()
	for i, v := range intarr2 {
		fmt.Printf("Value is %v at index %v\n", v, i)
	}

	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i)
		i++
	}

	var n int = 1000000
	var testslice = []int{}
	var testslice2 = make([]int, 0, n)

	fmt.Printf("\nTotal time without preallocation: %v\n", timeLoop(testslice, n))
	fmt.Printf("Total with with preallocation: %v\n", timeLoop(testslice2, n))

	var mystring = []rune("résumé")
	fmt.Println(len(mystring))
	for i, v := range mystring {
		fmt.Println(i, v)
	}

	var myrune = 'a'
	fmt.Printf("\nmyrun = %v, %T", myrune, myrune)

	var strslice = []string{"h", "i", " ", "t", "h", "ere"}
	var strbuilder strings.Builder
	for i := range strslice {
		strbuilder.WriteString(strslice[i])
	}
	var catstr = strbuilder.String()
	fmt.Println("\n" + catstr)
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
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