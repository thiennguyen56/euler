package main

import (
	"fmt"
	"math/big"
)
func get_sum(n int) int {
	number := big.NewInt(2)
	for i:=1; i < n; i++ {
		number = number.Mul(big.NewInt(2), number)
	}
	sum := 0
	for (number.Cmp(big.NewInt(0)) == 1) {
		modulo := big.NewInt(0).Mod(number, big.NewInt(10))
		sum += int(modulo.Int64())
		number = big.NewInt(0).Div(number.Sub(number, modulo), big.NewInt(10))
	}
	return sum

}

func main() {
	test := get_sum(1000)
	fmt.Println(test)
}