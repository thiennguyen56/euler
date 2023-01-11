package main

import (
	"fmt"
	"math"
)

func main() {
    x := 110000
    primes := make([]bool, x)
    p := 2
    for i := 0; i < x; i++ {
        primes[i] = true
    }
    for float64(p) <= math.Sqrt(float64(x)) {
        if primes[p] {
            for i:=p+p; i < x; i+=p {
                primes[i] = false
            } 
        }
        p += 1
    }
    primeNos := []int{}
    for i:=2; i < len(primes); i++ {
        if primes[i] {
            primeNos = append(primeNos, i)
        }
    }
    var T int
    fmt.Scan(&T)
    for t:=0; t < T; t++ {
        var N int
        fmt.Scan(&N)
        fmt.Println(primeNos[N-1])
    }
}