package main

import (
	"fmt"
)

func main() {
	i := 1
	var number float64

	fmt.Println("Enter Number: ")
	fmt.Scanf("%f", &number)

	for i <= int(number) {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else {
			fmt.Println(i)
		}
		i = i + 1
	}
}
