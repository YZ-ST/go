package main

import "fmt"

// 定義全域變數
var(
	a int = 10
	b string = "Hello"
)

// 定義常數
const (
    Unknown = 0
    Female = 1
    Male = 2
)


func main() {
	fmt.Printf("%d_%s", a, b)

	// print const unknown
	fmt.Printf("\n%d\n", Unknown)
}