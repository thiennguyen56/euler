package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Dog struct {
	Age int
}

func (d Dog) Eat() {
	fmt.Println("eat")
}

func (d Dog) Sleep(){
	fmt.Println("sleep")
}
func main(){
	testDog := Dog{}
	fmt.Println(testDog)
	testDog.Sleep()
	testDog.Eat()
}