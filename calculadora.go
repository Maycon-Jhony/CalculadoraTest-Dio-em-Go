package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculadora Simples em Go")

	var expressao string
	fmt.Print("Digite a expressão (por exemplo, 2 + 3 =): ")
	fmt.Scan(&expressao)

	resultado, err := calcular(expressao)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado)
}

func calcular(expressao string) (float64, error) {
	valores := strings.Split(expressao, " ")

	if len(valores) != 3 {
		return 0, fmt.Errorf("expressão inválida")
	}

	num1, err := strconv.ParseFloat(valores[0], 64)
	if err != nil {
		return 0, fmt.Errorf("valor inválido: %v", valores[0])
	}

	num2, err := strconv.ParseFloat(valores[2], 64)
	if err != nil {
		return 0, fmt.Errorf("valor inválido: %v", valores[2])
	}

	switch operacao := valores[1]; operacao {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("divisão por zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("operação desconhecida: %v", operacao)
	}
}
