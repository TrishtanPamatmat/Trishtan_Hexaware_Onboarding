package fizzbizz

import (
	"fmt"
)

// func main() {
// 	var number int
// 	fmt.Println("Enter numbers from 1 - 15: ")
// 	fmt.Scanf("%d", &number)

// 	Sample(number)
// }

func Sample(number int) []string {
	i := 1
	var arr []string
	for i <= int(number) {
		if i%3 == 0 {
			arr = append(arr, "Fizz")
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
			fmt.Println("Buzz")
		} else if i%15 == 0 {
			arr = append(arr, "Fizz Buzz")
			fmt.Println("Fizz Buzz")
		} else {
			arr = append(arr, fmt.Sprint(i))
			fmt.Println(i)
		}
		i = i + 1
	}
	return arr
}
