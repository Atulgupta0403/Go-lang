package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	// var ptr *int 

	// fmt.Println("Value of pointer is ", ptr) // ---nil---


	a := 5
	var ptr = &a


	fmt.Println("Value of pointer is ", ptr) // Address of a
	fmt.Println("Value of pointer is ", *ptr) // got the value on the address of a 


	*ptr = *ptr * 2
	fmt.Println("the main value of a is = ", a)  // a*2 = 10
	fmt.Println(*ptr)
	
	
}