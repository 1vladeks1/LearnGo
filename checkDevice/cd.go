package main

import "fmt"

type device struct{
	name string
	id string
	price float64
	isActive bool
}

func main() {
	d1 := device{
		name : "Принтер Canon",
		id : "45610311-88",
		price: 10900.99,
		isActive: true,
	}

	d2 := device{
		name : "Роутер Keenetic",
		id : "00190822-09",
		price: 3900.99,
		isActive: true,
	}

	d3 := device{
		name : "ПК-1/122",
		id : "11190781-92",
		price: 54000.99,
		isActive: false,
	}

	devices := []device{d1, d2, d3}

	var totalSum float64 = 0

	for _, d := range(devices){
		totalSum += d.price

		mode := "Устройство не работает"

		if d.isActive{
			mode = "Устройство работает"
		}
		
		fmt.Printf("Имя: %s;\nID: %s;\nЦена: %.2f рублей;\nСостояние: %s;\n\n",
		 			d.name, d.id, d.price, mode)
	}
	fmt.Println("Сумма всех устройств:", totalSum, "рублей")
}