package matematica

import (
	f "fmt"
	"strconv"
)

// Media calcula a média de números.
func Media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(f.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
