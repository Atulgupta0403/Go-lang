package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time Study of golang")

	presentaTime := time.Now()

	fmt.Println(presentaTime)

	fmt.Println("day-month-year")
	fmt.Println(presentaTime.Format("02-01-2006"))  // day-month-year
	
	fmt.Println("month-day-year")
	fmt.Println(presentaTime.Format("01-02-2006 15:04:05 Monday"))  // month-day-year  15:04:05--time

	// --------crreate time or date----------


	fmt.Println("--------crreate time or date----------")

	createdDate := time.Date(2004 , time.March , 4 , 12 , 14 , 0 ,0 , time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))

}