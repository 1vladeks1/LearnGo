package main

import "fmt"

func rps (p1, p2 string) string{
	if p1 == "paper" && p2 == "rock" ||
		p1 == "rock" && p2 == "scissors" ||
		p1 == "scissors" && p2 == "paper"{
			return "p1 wins!"
	} else if p1 == p2 {
		return "Draw!"
	} else {
		return "p2 wins!"
	}
}

func main(){
	a := rps("rock", "paper")
	fmt.Println(a)
}