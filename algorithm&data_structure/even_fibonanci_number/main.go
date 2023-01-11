package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	total := 0
	f1, f2 := 1, 2
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		for number > f2 {
			if f2%2 == 0 {
				total += f2
			}
			f1, f2 = f2, f1+f2
		}
		fmt.Println(total)
	}
}
