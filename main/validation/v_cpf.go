package validation

import (
	"strconv"
	"unicode"
)

// ValidaCPF verifica se um CPF é válido
func ValidaCPF(cpf string) bool {
	// Remove qualquer caractere que não seja número
	var numeros string
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			numeros += string(r)
		}
	}

	// CPF deve ter 11 dígitos
	if len(numeros) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais (ex: 111.111.111-11)
	todosIguais := true
	for i := 1; i < 11; i++ {
		if numeros[i] != numeros[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}

	// Calcula os dígitos verificadores
	digitos := numeros[:9]
	dv1 := calculaDigito(digitos, 10)
	dv2 := calculaDigito(digitos+dv1, 11)

	return numeros == digitos+dv1+dv2
}

// calculaDigito calcula um dígito verificador do CPF
func calculaDigito(digitos string, pesoInicial int) string {
	soma := 0
	for i, r := range digitos {
		num, _ := strconv.Atoi(string(r))
		soma += num * (pesoInicial - i)
	}
	resto := soma % 11
	if resto < 2 {
		return "0"
	}
	return strconv.Itoa(11 - resto)
}
