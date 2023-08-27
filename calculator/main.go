package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var action string
	fmt.Println("Enter action (+, -, *, /, %, sqrt): ")
	fmt.Scan(&action)
	if action != "+" && action != "-" && action != "*" && action != "/" && action != "%" && action != "sqrt" {
		fmt.Println("Invalid action")
		return
	}
	if action == "sqrt" {
		fmt.Println("Enter number: ")
		fmt.Scan(&a)
	} else {
		fmt.Println("Enter first number: ")
		fmt.Scan(&a)
	}
	if action != "sqrt" {
		fmt.Println("Enter second numbers: ")
		fmt.Scan(&b)
	}
	switch action {
	case "+":
		fmt.Printf("The sum of %v and %v is %v\n", a, b, a+b)
	case "-":
		fmt.Printf("The subtraction of %v and %v is %v\n", a, b, a-b)
	case "*":
		fmt.Printf("The multiplication of %v and %v is %v\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Printf("The division of %v and %v is %v\n", a, b, a/b)
	case "%":
		fmt.Printf("The remainder of %d and %d is %d\n", int(a), int(b), int(a)%int(b))
	case "sqrt":
		fmt.Printf("The square root of %v is %v\n", a, math.Sqrt(a))
	default:
		fmt.Println("Invalid action")
	}
}
