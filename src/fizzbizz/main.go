package main

import (
	"example_module/src/fizzbizz"
	"os"
	"strconv"
)

func main() {

	number := os.Args[1]
	intVar, _ := strconv.Atoi(number)

	fizzbizz.Sample(intVar)
}
