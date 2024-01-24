package main

import "testing"

func TestSoma(t *testing.T) {
	// Arrange
	teste, err := calcular("3 + 2")
	if err != nil {
		t.Fatal("Erro durante o cálculo:", err)
	}

	// Act
	resultado := 5.0

	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestCalcularSubtracao(t *testing.T) {
	// Arrange
	teste, err := calcular("10 - 4")
	if err != nil {
		t.Fatal("Erro durante o cálculo:", err)
	}

	// Act
	resultado := 6.0

	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestCalcularMultiplicacao(t *testing.T) {
	// Arrange
	teste, err := calcular("3 * 7")
	if err != nil {
		t.Fatal("Erro durante o cálculo:", err)
	}

	// Act
	resultado := 21.0

	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestCalcularDivisao(t *testing.T) {
	// Arrange
	teste, err := calcular("8 / 2")
	if err != nil {
		t.Fatal("Erro durante o cálculo:", err)
	}

	// Act
	resultado := 4.0

	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestCalcularDivisaoPorZero(t *testing.T) {
	// Arrange
	_, err := calcular("10 / 0")

	// Assert
	if err == nil || err.Error() != "divisão por zero" {
		t.Error("Erro esperado ao dividir por zero não ocorreu. Erro:", err)
	}
}
