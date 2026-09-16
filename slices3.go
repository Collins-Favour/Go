package main

import "fmt"

var x []int

func main() {
	x = []int{1, 2, 3}
	x = append(x, 10)
	fmt.Println(x)

}
