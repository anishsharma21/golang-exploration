package main

import "fmt"

var c, java, python = "no", true, false

func swap(x, y string) (string, string) {
	return y, x	
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var i, j int = 1, 2
	a, b := swap("there", "hi")
	k := 4
	fmt.Println(a, b)
	fmt.Println(split(18))
	fmt.Println(i, j, c, k, java, python)
}