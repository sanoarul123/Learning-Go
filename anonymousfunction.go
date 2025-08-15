package main

import "fmt"

// Standared or named function
func addTwo(x,y int){
	fmt.Println(x+y)
}

func anonymousFunction() {

	//Anonymous function havent any name
	// Immediately Invoked(Call kora) Function Expresession(value pass)

	//Interview Question: What is the difference between anonymous function and standard function?
								// What is IIFE
    func(x,y int){
		z:=x+y // its called expression
		fmt.Println("Hello i am anonymous function",z)
	}(5,6)
}

func init() {
	fmt.Println("Anonymous function example")

}
