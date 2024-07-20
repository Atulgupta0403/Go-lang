package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")

	var fruitList [5]string
	fruitList[0] = "gavava"
	fruitList[1] = "watermelon"
	fruitList[2] = "mango"
	fruitList[3] = "banana"
	fruitList[4] = "apple"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))


	var name = [3]string{"atul", "rahul" , "aditya"}
	fmt.Println(name)
	fmt.Println(len(name))
	
}