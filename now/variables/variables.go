package main

import "fmt"

func main() {

	firstNumber := 10
	secondNumber := 3.5

	sum1 := firstNumber + int(secondNumber)
	fmt.Println(sum1)

	sum2 := float64(firstNumber) + secondNumber
	fmt.Println(sum2)

	test := 10.4
	fmt.Println(test)

	test = 10
	fmt.Println(test)

}
