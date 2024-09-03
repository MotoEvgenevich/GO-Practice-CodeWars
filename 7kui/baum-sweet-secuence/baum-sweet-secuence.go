package kata

func BaumSweet(ch chan<- int) {
	defer close(ch) // Закрываем канал в конце

	for n := 0; n < 1000000; n++ {
		// Вычисляем значение Baum-Sweet для текущего n и отправляем его в канал

		ch <- baumSweetValue(n)
	}
}

// baumSweetValue вычисляет значение последовательности Баум-Суита для n
func baumSweetValue(n int) int {
	// Если n равно 0, то его двоичное представление считается пустой строкой, и результат 1
	if n == 0 {
		return 1
	}

	blockLength := 0
	for n > 0 {
		if n%2 == 0 { // Если текущий бит 0
			blockLength++
		} else { // Если текущий бит 1
			if blockLength%2 != 0 { // Если длина блока 0 нечетная
				return 0
			}
			blockLength = 0 // Сбрасываем длину блока после единицы
		}
		n /= 2
	}

	// Проверка последнего блока нулей, если таковой был
	if blockLength%2 != 0 {
		return 0
	}

	return 1
}
