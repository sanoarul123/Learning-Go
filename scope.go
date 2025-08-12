package main

import "fmt"

var (
	a = 20
	b = 30
)

func add(x int,y int){

	z := x+y
	fmt.Println(z)
}

//Todays Topic SCOPE
func scope() {
	//scope is very important in interview

	//Scope: should i access it? if yes. then there is scope. Scope is able to access 
	// var from anywhere

	p := 30
	q := 10
	
	add(p,q)
	add(a,b)
	add(a,z) // z isnt in the scope . because its declear in a function.





	fmt.Println("Hello")

}