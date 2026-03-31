// Анализ битовых ошибок

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s1 := []int{}
	s2 := []int{}

	var errors int

	for i := 0; i < 10; i++ {
		info := rand.Intn(2)
		s1 = append(s1, info)

		info = rand.Intn(2)
		s2 = append(s2, info)

		if s1[i] != s2[i] {
			errors++
		}
	}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("Количество ошибок: %d\n", errors)
}
