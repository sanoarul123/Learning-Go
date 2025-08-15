package main

import "fmt"


func functionExpressionAssignFunctionInVariable() {


	//Variable Expression example: a:=10
	/*
	if a<15 { //if block
		fmt.Println("Hello World")
	}
	*/
	//Function Assign into a varibale called functionExpression
	addition:= func(x,y int){
		z:=x+y // its called expression
		fmt.Println("Hello i am anonymous function",z)
	}

	addition(5,6)
}

func init() {
	fmt.Println("Function Expression And Assign Function In Variable example")

}
