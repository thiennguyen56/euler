package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	inputArray := [20][20]int{}
	i := 0
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		text := strings.Split(fileScanner.Text(), " ")
		for j := range text {
			iNumber, _ := strconv.Atoi(text[j])
			inputArray[i][j] = iNumber
		}
		i++
	}
	fmt.Println(inputArray)
	max := 1
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if (i < 17) {
				tmp := inputArray[i][j] * inputArray[i+1][j] * inputArray[i+2][j] * inputArray[i+3][j]
				if tmp > max {
					max = tmp
				}
			}
			if (j < 17) {
				tmp := inputArray[i][j] * inputArray[i][j+1] * inputArray[i][j+2] * inputArray[i][j+3]
				if tmp > max {
					max = tmp
				}
			}
			if (j < 17 && i < 17) {
				tmp := inputArray[i][j] * inputArray[i+1][j+1] * inputArray[i+2][j+2] * inputArray[i+3][j+3]
				if tmp > max {
					max = tmp
				}
			}
			if (j > 2 && i < 17) {
				tmp := inputArray[i][j] * inputArray[i+1][j-1] * inputArray[i+2][j-2] * inputArray[i+3][j-3]
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	fmt.Println(max)
}