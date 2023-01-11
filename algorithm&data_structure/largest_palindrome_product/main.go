package main

import (
	"fmt"
	"strconv"
)

func get_result(n int) { 
    for i := n - 1; i >= 101101; i-- {
        status := 1
        nStr := strconv.Itoa(i)
        check:
        for x, y := 0, 5; x < 3; x, y = x+1, y-1 {
            if nStr[x] != nStr[y]{
                status = 0
                break check
            }
        }
        
        if status == 1 {
            for j := 100; j < 1000; j++ {
                quotient := i /j
                remainder := i % j
                if remainder == 0 && quotient >= 100 && quotient <= 999 {
                    fmt.Println(i)
                    return
                }
            }
        }
    }
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var T int
    fmt.Scan(&T)
    for i := 0; i < T; i++ {
        var N int
        fmt.Scan(&N)
        get_result(N)
    }
}