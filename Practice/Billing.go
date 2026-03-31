// Биллинг данных

package main

import (
	"fmt"
	//"math/rand"
)

type abonent struct {
	name      string
	ratePrice float64
	usage     float64
	isActive  bool
}

func main() {
	var totalPrice float64

	debtors := []string{}

	abonents := map[string]abonent{
		"+7 989 712 19 07": {
			"Мрыхин ВЮ", 800.99, 26.78, true,
		},

		"+7 800 555 35 35": {
			"Иванов ИИ", 553.00, 5.67, false,
		},

		"+7 999 165 94 01": {
			"Петров ОА", 450.20, 0.96, true,
		},
	}

	for user, value := range abonents {
		status := "Неактивный"
		if value.isActive {
			status = "Активный"
		}

		totalPrice += value.ratePrice

		if abonents[user].usage > 5.0 {
			debtors = append(debtors, user)
		}

		fmt.Printf("%s - Имя: %s; Цена тарифа: %.2f рублей; Использовано: %.2f Гб; Активность: %s\n",
			user, value.name, value.ratePrice, value.usage, status)
	}
	fmt.Printf("\nОбщая стоимость всех тарифов: %.2f рублей\n\n", totalPrice)

	for _, name := range debtors {
		fmt.Printf("Должник: %s, %.2f Гб\n", abonents[name].name, abonents[name].usage)
	}
}
