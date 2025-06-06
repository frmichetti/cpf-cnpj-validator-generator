package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// GeraCPF cria um CPF válido (com os dois dígitos verificadores)
func GeraCPF() string {
	rand.Seed(time.Now().UnixNano())

	// Gera os 9 primeiros dígitos aleatórios
	base := make([]int, 9)
	for i := 0; i < 9; i++ {
		base[i] = rand.Intn(10)
	}

	// Calcula o primeiro dígito verificador
	dv1 := calculaDigito1(base, 10)

	// Adiciona o primeiro DV e calcula o segundo
	base = append(base, dv1)
	dv2 := calculaDigito1(base, 11)

	// Monta CPF completo
	base = append(base, dv2)

	// Formata como string com separadores
	return fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d",
		base[0], base[1], base[2],
		base[3], base[4], base[5],
		base[6], base[7], base[8],
		base[9], base[10],
	)
}

// calculaDigito calcula um dos dígitos verificadores do CPF
func calculaDigito1(nums []int, pesoInicial int) int {
	soma := 0
	for i := 0; i < len(nums); i++ {
		soma += nums[i] * (pesoInicial - i)
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
