package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"apple","banana","orange"}
	fmt.Println(fruitList) //type string hoga iska 
	
	fruitList = append(fruitList, "grapes","gauava")
	fmt.Println(fruitList) //type string hoga iska 
	
	fruitList = append(fruitList[2:3])
	fmt.Println(fruitList)


	highScroes := make([]int, 5)

	highScroes[0] = 523
	highScroes[1] = 223
	highScroes[2] = 333
	highScroes[3] = 133
	highScroes[4] = 033
	// highScroes[5] = 533
	fmt.Println(highScroes)
	fmt.Println(len(highScroes))


	// highScroes = append(highScroes,666, 777,888,999)
	// fmt.Println(highScroes)
	// fmt.Println(len(highScroes))

	fmt.Println(sort.IntsAreSorted(highScroes))
	sort.Ints(highScroes)
	fmt.Println(highScroes) //it is only present in slices not in array , slices are heavily user in go lang 
}