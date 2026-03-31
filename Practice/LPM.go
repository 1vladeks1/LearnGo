// Таблица маршрутизации

package main

import "fmt"

func main() {
	devices := map[string]string{}

	var device_ip, device_name string

	for i := 0; i < 5; i++ {
		fmt.Print("Введите IP-адрес и название устройства через пробел: ")
		fmt.Scanln(&device_ip, &device_name)

		devices[device_ip] = device_name
	}
	fmt.Println()

	fmt.Println("Список устройств:")
	for ip, name := range(devices){
		fmt.Printf("\t%s : %s\n", ip, name)
	}

	var check string

	fmt.Print("Введите IP-адрес для проверки: ")
	fmt.Scanln(&check)

	if search, ok := devices[check]; ok{
		search = "Устройство есть в таблице"
		fmt.Println(search)
	} else {
		search = "Устройства нет в таблцие"
		fmt.Println(search)
	}
	
}
