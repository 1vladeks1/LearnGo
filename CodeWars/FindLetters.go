package main

import "fmt"

func GetCount(str string) (count int){
	count = 0

	for _, l := range str{
		switch l{
			case 'a', 'o', 'e', 'i', 'u':
				count++
			default:
				continue
		}
	}
	return count
}

func main(){
	a := GetCount("abracadabra")
	fmt.Println(a)
}