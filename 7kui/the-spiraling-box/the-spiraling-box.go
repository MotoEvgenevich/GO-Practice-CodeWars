package kata

func CreateBox(width, length int) [][]int {
	// Инициализируем двумерный срез с нулями
	matrix := make([][]int, length)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	// Определяем количество слоев
	layers := (Min(width, length) + 1) / 2
	for layer := 0; layer < layers; layer++ {
		value := layer + 1

		// Заполняем верхнюю и нижнюю строки для текущего слоя
		for i := layer; i < width-layer; i++ {
			matrix[layer][i] = value          // Верхняя строка
			matrix[length-layer-1][i] = value // Нижняя строка
		}

		// Заполняем левый и правый столбцы для текущего слоя
		for i := layer; i < length-layer; i++ {
			matrix[i][layer] = value         // Левый столбец
			matrix[i][width-layer-1] = value // Правый столбец
		}
	}

	return matrix
}

// Функция для нахождения минимального значения между двумя числами потому что Codewars использует старую версию Go и там ее нет!!!
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
