package main

import "fmt"

// a এগে নাহ p আগে মনে রাখার কৌশল
func summationAdd(a int, b int) int { // receiving value is called parameter: a, b
	return a + b
}

func pstsVsArg(a,b int)  {
	sum:=summationAdd(a, b)
	fmt.Println("Sum is",sum)
	
}

func init() {
	fmt.Println("I am in pstsVsArg.go")
}