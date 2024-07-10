package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("rate my app")
	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')
	fmt.Println("your rating is " , input)

	// numRating , err := strconv.ParseFloat(input , 64)  -----error := kuy ki input me too "\n" v tha n too woo string too hai nhi isi liye error aa rha tha ----
	numRating , err := strconv.ParseFloat(strings.TrimSpace(input) , 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating : ", numRating +1)
	}


	
}