package main

import (
	"fmt"

	"github.com/snehil-sinha/GoTesting/math"
)

func main() {
	var number int
	fmt.Print("Enter the number: ")
	fmt.Scan(&number)
	fmt.Println("The absolute value of the entered number is: ", math.Abs(number))
}
