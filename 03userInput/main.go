package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter any string")

	// comma ok / err err
	input , _ := reader.ReadString('\n')
	fmt.Println("thanks to write " , input)
	
}