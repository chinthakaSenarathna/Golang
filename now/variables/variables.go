package main

import "fmt"

const pi float64 = 3.14

func main() {

	var radius float64

	fmt.Print("Input the radius : ")
	_, err := fmt.Scan(&radius)

	if err != nil {
		fmt.Print("Invalid radius")
		return
	}

	area := pi * radius * radius
	fmt.Printf("Area is %.2f", area)

}
