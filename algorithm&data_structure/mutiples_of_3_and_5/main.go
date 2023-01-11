package main

import (
	"fmt"
)

func sum_unit(n int) int {
	return n*(n+1) / 2
}

func main(){
	var n int
	fmt.Scan(&n)
	for i:=0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		number -= 1
		fmt.Println(3 * sum_unit(number / 3) + 5 * sum_unit(number / 5) - 15 * sum_unit(number / 15))
	}
}