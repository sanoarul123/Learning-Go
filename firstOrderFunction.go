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

	Higher Order Function- any of the following three-
		1. parameter -> function
		2. function return
		3. both

*/

func processOperation(a,b int,op func (p,q int)){

	op(a,b)
}




func addition(x,y int){
	z := x + y
	fmt.Println("Hi,itz me",z)
}

func firstOrderFunction(){
	processOperation(10,20,addition) 


}

func init() {	

	fmt.Println("First Order Function example")
}