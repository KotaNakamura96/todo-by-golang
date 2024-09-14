package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 変数名も指定することができる.その場合はどの変数をreturnするかは省略できる.
func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return.")
}

func main() {
	i1 := Plus(1, 4)
	fmt.Println(i1)

	i2, _ := Div(9, 4)
	fmt.Println(i2)

	i4 := Double(10)
	fmt.Println(i4)

	Noreturn()
}