package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	other_ints := integers()
	fmt.Println(other_ints())
	fmt.Println(other_ints())
	fmt.Println(other_ints())
}