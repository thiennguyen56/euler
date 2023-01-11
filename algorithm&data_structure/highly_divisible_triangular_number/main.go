package main

import "fmt"

func get_list_factor(n int) []int {
	factors := []int{}
	for i:=1; i <= n; i++ {
		if n % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
func main(){
	step := 1

	start := 0
	terms := []int{}
	list_factor := []int{}
	for len(list_factor) <= 500 {
		terms = append(terms, start+step)
		list_factor = get_list_factor(start+step)
		step += len(terms) + 1
	}
	fmt.Println(terms[len(terms) - 1])
	// fmt.Println(list_factor)

	
}