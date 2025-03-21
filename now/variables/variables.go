package main

import "fmt"

func main() {
	var firstNumber int = 100
	var secondNumber int = 11

	fmt.Println(firstNumber + secondNumber)
	fmt.Println(firstNumber - secondNumber)
	fmt.Println(firstNumber * secondNumber)
	fmt.Println(firstNumber / secondNumber)
	fmt.Println(firstNumber % secondNumber)

	var userName string = "chinthaka senarathna"

	fmt.Println("hello... " + userName)

	var firstNumber_ float32 = 1.2
	var secondNumber_ float64 = 1.23

	fmt.Printf("Type : %T\n", firstNumber_)
	fmt.Printf("Type : %T\n", secondNumber_)

	var isUser bool = true
	if isUser {
		println("This is the User")
	}

	var firstName, secondName = "chinthaka", "senarathna"
	fmt.Println(firstName, secondName)

}
