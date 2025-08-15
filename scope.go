package main

import "fmt"

var (
	a = 20
	b = 30
)

//Todays Topic SCOPE
func scope() {
	//scope is very important in interview

	//Scope: should i access it? if yes. then there is scope. Scope is able to access 
	// var from anywhere

	//scope is three types: global, local, package
	//Local scope types: if , function, switch 
	//block : block means curly brase {}.
	//inside block: All variable called local variable

	p := 30
	q := 10
	
	add(p,q)
	add(a,b)
	add(a,p) // z isnt in the scope . because its declear in a function.

	if(p>20){
		var a = 31
		fmt.Println("You are",a)
	}
	fmt.Print(a)





	fmt.Println("Hello")

}