package services

func CalcularNotas(valor int) (map[int]int, bool) {
	notas := []int{200, 100, 50, 20, 10, 5}
	resultado := make(map[int]int)

	for _, nota := range notas {
		if valor >= nota {
			resultado[nota] = valor / nota
			valor = valor % nota
		}
	}

	if valor != 0 {
		return nil, false
	}

	return resultado, true
}
