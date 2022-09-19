package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y int

	fmt.Println("Karatsuba Multiplication")
	fmt.Print("Introduce Multiplicand: ")
	fmt.Scan(&x) // no value type validation
	fmt.Print("Introduce Multiplier: ")
	fmt.Scan(&y) // no value type validation
	fmt.Println("Product: ", karatsuba(x, y))
}

func karatsuba(x, y int) int {
	// base case
	if x < 10 && y < 10 {
		return x * y
	}

	// calculate the size of the numbers
	x_len := float64(len(strconv.Itoa(x)))
	y_len := float64(len(strconv.Itoa(y)))

	/*
		Divide the numbers in two halves. For example:

		x: 1234 , then a: 12 ; b: 34
		y: 567, then c: 56; d: 7

		In this example n value is: 2 and m value is 100
	*/
	n := int(math.Max(x_len, y_len)) / 2
	m := int(math.Pow10(n))

	a := int(x / m)
	b := x % m

	c := int(y / m)
	d := y % m

	// make recursive calls
	ac := karatsuba(a, c)
	bd := karatsuba(b, d)
	e := karatsuba(a, d) + karatsuba(b, c)

	return (int(math.Pow10(n*2)) * ac) + (int(math.Pow10(n)) * e) + bd
}
