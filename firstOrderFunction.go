package main

import "fmt"

/*
	first Order Function. -- influenced by Functional Programing language
		1. Starndeard function or named function
		2. anonymous function
		3. IIFE
	এইগুলাই বুঝানো হয়।

	functinal paradigm -> haskel,racket,

	math -> logic (discreate mathematic)
		1. first order logic - is working based on Object,Property and Relation
		2. higher order logic - is woking with rule


	logic

	1. Object (people,animal,car etc)
	2. Property(color,student)
	3.Relation


	Tutul is a student
	Apple is red
	Tutul is  taller than rakib

	Statement


	Example :
		Rule: All customers Must pay their pizza bills
		Fact: Tutul is a customer
		Conclusion: Tutul must pay his pizza bill

		again :
		rule: All customers must pay tips to the waiters.

*/

/*

	Higher Order Function- any of the following three- higher order function is also called first class function
		1. parameter -> function
		2. function return
		3. both

*/
/*

	Call Back function - which function pass into higher order function as argument


*/

/*

first class citizen - a programming language feature that treats functions as first-class citizens,
							allowing them to be assigned to variables, passed as arguments, and returned from other functions.


*/
func processOperation(a,b int,op func (p,q int)){ // receive as parameter function called higher order function

	op(a,b)
}

func call() func (x,y int) { // return function is also called higher order function
	return	addition
}

func addition(x,y int){
	z := x + y
	fmt.Println("Hi,itz me",z)
}

func firstOrderFunction(){
	// processOperation(10,20,addition) 
		sum := call()
		sum(5,6)


}

func init() {	

	fmt.Println("First Order Function example")
}