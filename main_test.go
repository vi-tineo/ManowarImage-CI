package main

import "testing"

func TestCalcularPreco(t *testing.T) {
	total := calcularPreco(2, true)
	if total != 1000.0 {
		t.Errorf("Esperado 1000.0, obtido %.2f", total)
	}
}
