package main

import (
	"fmt"
	"math/rand"
)

func main() {
	server := map[string] []int{
		"192.168.1.1" : {},
		"127.0.0.1" : {},
		"96.45.16.11" : {},
	}	

	for i := range server{
		for j := 0; j < 5; j++{
			min, max := 30, 400
			ping := rand.Intn(max - min + 1) + min
			server[i] = append(server[i], ping)
		}
	}

	for id, pings := range server{
		var midPing = 0

		
		fmt.Printf("Пинг для %11s:\n", id)
		
		for p := 0; p < len(pings); p++{
			midPing += pings[p]
			if pings[p] > 100{
				fmt.Printf("\t - %v (ВЫСОКАЯ ЗАДЕРЖКА)\n", pings[p])
				continue
			}
			fmt.Printf("\t - %v\n", pings[p])

		}
		fmt.Println()

		fmt.Printf("Средний пинг для %s:\n", id)
		fmt.Printf("\t-%d\n", midPing / len(pings))
		fmt.Printf("---------------------------------------\n")
	}
}