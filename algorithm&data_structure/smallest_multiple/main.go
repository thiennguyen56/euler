package main

import (
	"fmt"
	"math"
)
func getArrayPrime(n int) []int {
	i, maxi := 2, 0
	result := []int{}
	for i*i <= n {
		for (n % i) == 0 {
			n /= i
			maxi = i
			result = append(result, maxi)
		}
		i++
	}
	if n > maxi {
		result = append(result, n)
	}
	return result
}

func get_result(n int) int {
	mapPrime := make(map[int]int)
	for i := 1; i <= n; i++ {
		arrayPrime := getArrayPrime(i)
		counter := make(map[int]int)
		for _, val := range arrayPrime {
			counter[val]++
		}
		for k, v := range counter {
			if _, ok := mapPrime[k]; !ok {
				mapPrime[k] = v
			} else if v > mapPrime[k] {
				mapPrime[k] = v
			} 
		}
	}
	total := 1
	for k, v := range mapPrime {
		total *= int(math.Pow(float64(k), float64(v)))
	}
	return total
}

func main(){
	var n int
	fmt.Scan(&n)
	for i:=0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		fmt.Println(get_result(number))
	}
}