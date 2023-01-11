package main

import "fmt"

func main() {
    var N, n int
    fmt.Scan(&N)
    for i:=0; i < N; i++ {
        fmt.Scan(&n)
        j := 2
        maxj := 0
        for j*j < n {
            for n % j == 0 {
                maxj = j
                n /= j
            }
            j++
        }
        if (n > maxj) {
            maxj = n 
        }
        fmt.Println(maxj)
    }

}