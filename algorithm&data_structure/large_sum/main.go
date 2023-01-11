package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	readFile, err := os.Open("text.txt")

	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	total := new(big.Int)
	for fileScanner.Scan(){
		n := new(big.Int)
		n, _ = n.SetString(fileScanner.Text(), 10)
		total.Add(total, n)
	}
	fmt.Println(total)
	
}