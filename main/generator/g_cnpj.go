package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// GeraCNPJ retorna um CNPJ válido no formato 00.000.000/0000-00
func GeraCNPJ() string {
	rand.Seed(time.Now().UnixNano())

	// Gera os 8 primeiros dígitos aleatórios (raiz do CNPJ)
	base := make([]int, 8)
	for i := 0; i < 8; i++ {
		base[i] = rand.Intn(10)
	}

	// Fixa a parte do número do estabelecimento como "0001"
	base = append(base, 0, 0, 0, 1)

	// Calcula os dois dígitos verificadores
	dv1 := calculaDigitoCNPJ1(base, []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	base = append(base, dv1)
	dv2 := calculaDigitoCNPJ1(base, []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	base = append(base, dv2)

	// Formata como string
	return fmt.Sprintf("%d%d.%d%d%d.%d%d%d/%d%d%d%d-%d%d",
		base[0], base[1], base[2], base[3], base[4],
		base[5], base[6], base[7], base[8], base[9],
		base[10], base[11], base[12], base[13],
	)
}

// calculaDigitoCNPJ calcula um dígito verificador do CNPJ
func calculaDigitoCNPJ1(nums []int, pesos []int) int {
	soma := 0
	for i := 0; i < len(pesos); i++ {
		soma += nums[i] * pesos[i]
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
