package main

import (
	"fmt"
	"math/rand"
	"time"
)

type user struct{
	ID int
	Balance float64
	isActive bool
}

func (u *user) deposit (money float64){
	u.isActive = true
	u.Balance -= money
	fmt.Printf("Списано %.1f рублей со счёта.\n", money)
	fmt.Printf("Состояние активности: %t\n", u.isActive)
	fmt.Printf("На вашем счету: %.2f\n\n", u.Balance)
	time.Sleep(1 * time.Second)
}

func (u *user) check (money float64) error{
	if u.Balance < money {
		u.isActive = false
		err := fmt.Errorf("Недостаточно средств!\n")
		fmt.Println(err)
	}
	return nil
}

func main(){
	me := user{
		ID : 19089032,
		Balance: 560.35,
		isActive: false,
	}

	for i := 0; i < 5; i++{
		min, max := 50, 200
		me.deposit(float64(rand.Intn(max - min + 1) + min))
		me.check(float64(rand.Intn(max - min + 1) + min))
	}
}