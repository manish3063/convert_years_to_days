package main

import "fmt"

func main() {
	days := 373

	year := days / 365
	reminder := days % 365
	month := reminder / 30
	mreminder := reminder % 30
	week := mreminder / 7
	rem_week := mreminder % 7
	day := rem_week
	fmt.Println("No of year", year)
	//fmt.Println(reminder)
	fmt.Println("No of month", month)
	//fmt.Println(mreminder)
	fmt.Println("No of weeks", week)
	//fmt.Println(rem_week)
	fmt.Println("No of days", day)
}
