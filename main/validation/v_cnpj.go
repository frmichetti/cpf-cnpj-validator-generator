package validation

import (
	"strconv"
	"unicode"
)

// ValidaCNPJ valida se um CNPJ é válido
func ValidaCNPJ(cnpj string) bool {
	// Remove qualquer caractere não numérico
	var numeros string
	for _, r := range cnpj {
		if unicode.IsDigit(r) {
			numeros += string(r)
		}
	}

	if len(numeros) != 14 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	todosIguais := true
	for i := 1; i < 14; i++ {
		if numeros[i] != numeros[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}

	// Cálculo dos dígitos verificadores
	base := numeros[:12]
	dv1 := calculaDigitoCNPJ(base, []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	baseComDv1 := base + dv1
	dv2 := calculaDigitoCNPJ(baseComDv1, []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	return numeros == base+dv1+dv2
}

// calculaDigitoCNPJ calcula um dígito verificador do CNPJ com pesos fornecidos
func calculaDigitoCNPJ(base string, pesos []int) string {
	soma := 0
	for i, r := range base {
		num, _ := strconv.Atoi(string(r))
		soma += num * pesos[i]
	}
	resto := soma % 11
	if resto < 2 {
		return "0"
	}
	return strconv.Itoa(11 - resto)
}
