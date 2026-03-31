// Мониторинг сети

package main

import "fmt"

type network struct {
	name string
	capacity float64
	usage float64
}

func main() {
	devices := []network{
		{"TP-LINK-MB89", 20.45, 11.5},
		{"KEENETIC 1121", 11.1, 2.41},
		{"XIAOMI-4AC", 100.56, 6.89},
		{"TP-LINK-881K", 20.11, 18.2},
	}

	var netUsage float64 = 0

	for _, info := range devices{
		netUsage = (info.usage / info.capacity) * 100

		fmt.Printf("\nНагрузка на сеть устройства %s: %.2f%%.\n", info.name, netUsage)

		if netUsage > 50 && netUsage < 90{
			fmt.Println("\t- ВНИМАНИЕ, устройство начинает нагружаться")
		} else if netUsage > 90{
			fmt.Println("\t- !!!ВНИМАНИЕ!!!, устройство перегружено")
		} else {
			fmt.Printf("\t- Устройство работает стабильно\n")
		}

		
	}
}
