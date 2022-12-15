package main

import (
	"fmt"
)

func main() {
	var codigo, quantidade int
	fmt.Scan(&codigo)
	fmt.Scan(&quantidade)
	total := 0.0

	switch codigo {
	case 1:
		total = float64(quantidade) * 4.00
		break
	case 2:
		total = float64(quantidade) * 4.50
		break
	case 3:
		total = float64(quantidade) * 5.00
		break
	case 4:
		total = float64(quantidade) * 2.00
		break
	case 5:
		total = float64(quantidade) * 1.50
		break
	default:
		total = 0
	}

	fmt.Printf("Total: R$ %.2f\n", total)
}
