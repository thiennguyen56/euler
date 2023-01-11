package main

import "fmt"

func collazSequence(n int) (int, int) {
	count := 0
	number := n
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
		count++
	}
	return count, number
}

func main() {
	number := 0
	max := 0
	for i := 1; i < 1000000; i++ {
		tmp, n := collazSequence(i)
		if tmp > max {
			number = n
			max = tmp
		}
	}
	fmt.Println(number)
}