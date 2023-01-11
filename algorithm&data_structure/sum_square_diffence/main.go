package main

import (
	"fmt"
	"math"
)

func get_result(n int) int {
	sum_square_n := (n * (n + 1) * (2*n + 1)) / 6
	square_sum_n := int(math.Pow(float64(n * (n + 1) / 2), 2))
	return square_sum_n - sum_square_n
}

func main(){
	var T int
	fmt.Scan(&T)
	for i:=0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		fmt.Println(get_result(N))
	}
}