package main

import "fmt"

func smallestInt (numbers []int) int {
	min := numbers[0]

	for _, value := range numbers{
		if value < min{
			min = value
		} 
	}
	return min
}

func main(){
	a := smallestInt([]int{34, 15, 88, 2})
	fmt.Println(a)
}