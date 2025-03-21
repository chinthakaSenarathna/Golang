package main

import "fmt"

func main() {

	firstNumber := 10
	secondNumber := 3.53

	ans1 := firstNumber * int(secondNumber)
	fmt.Println(ans1)
	ans2 := float64(firstNumber) * secondNumber
	fmt.Println(ans2)

	ans3 := firstNumber / int(secondNumber)
	fmt.Println(ans3)

	ans4 := float64(firstNumber) / secondNumber
	fmt.Println(ans4)

}
