package main

import (
	"fmt"
	"math"
)

func main() {
    var T int
    fmt.Scan(&T)
    for i := 0; i < T; i++ {
        var S int
        m := -1
        fmt.Scan(&S)
        for a:=3; a < (S/3) + 1; a++ {
            b := (int(math.Pow(float64(S), 2.0)) - 2*a*S) / (2*S - 2*a)
            c := S - a -b 
            if a*a + b*b == c*c && a*b*c > m {
                m = a*b*c
            }
        }
        fmt.Println(m)
    }
}