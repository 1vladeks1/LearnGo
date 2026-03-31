// Мощность сигнала

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	Ptx := rand.Intn(20) + 5  
	Gtx := rand.Intn(20) + 5  
	Grx := rand.Intn(20) + 5  
	d := rand.Intn(100)
	Lambda := rand.Float64() + 0.2

	loss := math.Pow(Lambda / (4 * math.Pi * float64(d)), 2)

	Prx := float64(Ptx) * float64(Gtx) * float64(Grx) * loss
	fmt.Printf("Мощность сигнала: %.13f Вт\n", Prx)
	// Prx << Ptx!!! - Закон сохранения энергии

	Prx_dB := 10 * math.Log10(Prx/0.001)
	fmt.Printf("Мощность сигнала: %.2f дБм\n", Prx_dB)
}
