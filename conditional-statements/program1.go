package main

import (
	"fmt"
	"strings"
	"unicode"
)

func parity(num int) string {
	if num < 0 {
		return "negative"
	} else {
		return "positive"
	}
}

func dayOfWeek(day int) string {
	if day > 5 {
		return "Weekend"
	}
	return "Weekday"
}

func isConsonant(c rune) bool {
	c = unicode.ToLower(c)
	return unicode.IsLetter(c) && !strings.ContainsRune("aeiou", c)
}

func factorial(n int) int {
	value := 1
	for i := 1; i <= n; i ++ {
		value *= i
	}
	return value
}

func factorialRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println(parity(3))
	fmt.Println(dayOfWeek(3))
	fmt.Println(dayOfWeek(7))
	fmt.Println(isConsonant('t'))
	fmt.Println(isConsonant('a'))
	fmt.Println(factorial(5))
	fmt.Println(factorial(12))
	fmt.Println(factorialRecursive(5))
}