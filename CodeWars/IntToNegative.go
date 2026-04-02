package main

import "fmt"

func makeNegative (x int) int {
	var negative int = 0

	if x > 0{
		negative = x * -1
	} else if x < 0 {
		negative = x
	} else if x == 0 {
		negative = 0
	}
	return negative
}

func main(){
	a := makeNegative(1)
	b := makeNegative(-10)
	c := makeNegative(0)
	fmt.Println(a, b, c)
}