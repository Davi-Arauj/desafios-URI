package main

import (
	"fmt"
)

func main() {
	valor := 0.0
	fmt.Scanln(&valor)
	mensagemTrue := ""

	if valor >= 0 && valor <= 25 {
		mensagemTrue = "Intervalo [0,25]"
	} else if valor > 25 && valor <= 50 {
		mensagemTrue = "Intervalo (25,50]"
	} else if valor > 50 && valor <= 75 {
		mensagemTrue = "Intervalo (50,75]"
	} else if valor > 75 && valor <= 100 {
		mensagemTrue = "Intervalo (75,100]"
	} else {
		mensagemTrue = "Fora de intervalo"
	}

	fmt.Println(mensagemTrue)
}
