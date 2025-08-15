package main

/*Topic:

1. Variable Shadow 019 episod
2. Standared Function: যার নাম আছে বা যেই ফাংকশন এর নাম আচে তাকে বলে
3. Init Function: যাকে আমি কল করতে পারব নাহ, কম্পিউটার নিজে থেকেই কল করে
		i. No input / no output / no return










*/

//to use custom package first create this by : go mod init example.com

import (
	"fmt"

	"example.com/mathlib"
)

func unknown2(temp1 int,temp2 int,temp3 string) (int,int){

	if temp3 == "sum" {
		sum := temp1 + temp2
		fmt.Println("Sumation Is",sum)
		mul := 20
		return sum,mul 
	} else if temp3 == "multiple"{
		multiple := temp1* temp2
		fmt.Println("Multiple Is",multiple)
		mul := 20
		return multiple,mul 
	} else if temp3 == "division"{
		divition := temp1/temp2
		fmt.Println("Divition Is",divition)
		mul := 20
		return divition ,mul


	}
	return 0,0
}


func unknown(temp1 int,temp2 int,temp3 string){

	if temp3 == "sum" {
		sum := temp1 + temp2
		fmt.Println("Sumation Is",sum)
	} else if temp3 == "multiple"{
		multiple := temp1* temp2
		fmt.Println("Multiple Is",multiple)
	} else if temp3 == "division"{
		divition := temp1/temp2
		fmt.Println("Divition Is",divition)
	}
}


func add(temp1 int,temp2 int) {

	sum := temp1 + temp2

	fmt.Println("Sumation Is",sum)

}

func main() {

	fmt.Println("Hello World")

	//Variable Declaration

	// var x int = 10

	// var x = "SANOAR"

	x:=10

	fmt.Println("Hello ",x)

	//if else

	age := 18

	if age > 18 {
		fmt.Println("Hi. You are above 18")
	} else if age < 18 {
		fmt.Println("Hi. You are under 18")
	} else {
		fmt.Println("Hi. You are Equal 18")
	}

	//switch

	switch age {
	case 1:
		fmt.Println("Hi")
	case 2:
		fmt.Println("Hello")
	default:
		fmt.Println("Sorry")
	}

	//function

	add(5,6)
	unknown(5,6,"multiple")

	//take input from user

	age1,age2 := unknown2(5,6,"multiple")
	fmt.Println(age1,age2)

	var firstNumber ,secondNumber int
	var typeOf string
	fmt.Println("Enter First & Second Number Respectively:") 
	fmt.Scanln(&firstNumber,&secondNumber)

	fmt.Println("What you want:")
	fmt.Scanln(&typeOf)
	unknown(firstNumber,secondNumber,typeOf)
	scope() 

	
	//end 



	fmt.Println("Showing Custom Package")
	mathlib.Add(5,6)




	





}

func init(){
	fmt.Println("Initialization")
}