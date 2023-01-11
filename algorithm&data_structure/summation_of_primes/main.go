package main

import (
	"fmt"
)
func main(){
	var n int
	fmt.Scan(&n)
	numbers := []int{}
	for i:=0; i <=n; i++ {
		numbers = append(numbers, i)
	}

	for i:=2; i <= n/2; i++ {
		if numbers[i] != 0 {
			for j:= i*i; j <= n; j+=i {
				numbers[j] = 0
			}
		}
	}
	total := 0
	for i:=2; i <=n; i++ {
		total += numbers[i]
	}
	fmt.Println(total)
}