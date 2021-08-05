package main

import "fmt"

func main() {
	i := 1
	for i <= 15 {
		if i%3 == 3 {
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
