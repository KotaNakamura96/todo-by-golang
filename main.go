package main

import "fmt"

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true || true == false)
}