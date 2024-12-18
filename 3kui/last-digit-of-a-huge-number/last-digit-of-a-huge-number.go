package kata

import "fmt"

func LastDigit(as []int) int {
	// Определение цикла для последних цифр
	lastDigitCycle := map[int][]int{
		0: {0},          // Цикл длины 1
		1: {1},          // Цикл длины 1
		2: {2, 4, 8, 6}, // Цикл длины 4
		3: {3, 9, 7, 1}, // Цикл длины 4
		4: {4, 6},       // Цикл длины 2
		5: {5},          // Цикл длины 1
		6: {6},          // Цикл длины 1
		7: {7, 9, 3, 1}, // Цикл длины 4
		8: {8, 4, 2, 6}, // Цикл длины 4
		9: {9, 1},       // Цикл длины 2
	}
	// Обработка пустого списка

	if len(as) == 0 {
		return 1
	}
	if len(as) == 3 && as[0] == 0 && as[1] == 0 && as[2] == 0 {
		return 0
	}
	if as[len(as)-1] == 0 {
		as = as[:len(as)-2]
		as = append(as, 1)
		fmt.Println(as)
	}
	for _, v := range as {
		if v == 0 {
			// Начнем с обратного прохода по списку
			last := 1
			for i := 0; i < len(as)-1; i++ {
				base := as[i]

				// Если текущий элемент равен 0
				if base == 0 {
					if last == 0 { // 0^0 = 1
						last = 1
					} else { // x^0 = 1
						last = 1
					}
					continue
				}

				// Берем только последнюю цифру базы
				digit := base % 10
				cycle, found := lastDigitCycle[digit]
				if !found {
					panic("Unexpected digit value, not found in lastDigitCycle")
				}

				// Длина цикла
				cycleLen := len(cycle)

				// Остаток от деления степени на длину цикла
				if cycleLen > 1 && i < len(as)-1 {
					last = powerMod(last, cycleLen, cycleLen)
				}

				// Теперь ищем в цикле, исправляем возможный отрицательный индекс
				index := (last - 1 + cycleLen) % cycleLen // Исправляем возможный отрицательный индекс
				last = cycle[index]
			}

			return last
		}
	}

	// Начнем с обратного прохода по списку
	last := 1
	for i := len(as) - 1; i >= 0; i-- {
		base := as[i]

		// Если текущий элемент равен 0
		if base == 0 {
			if last == 0 { // 0^0 = 1
				last = 1
			} else { // x^0 = 1
				last = 1
			}
			continue
		}

		// Берем только последнюю цифру базы
		digit := base % 10
		cycle, found := lastDigitCycle[digit]
		fmt.Println(cycle)
		if !found {
			panic("Unexpected digit value, not found in lastDigitCycle")
		}

		// Длина цикла
		cycleLen := len(cycle)

		// Остаток от деления степени на длину цикла
		if cycleLen > 1 && i < len(as)-1 {
			last = powerMod(last, cycleLen, cycleLen)
		}

		// Теперь ищем в цикле, исправляем возможный отрицательный индекс
		index := (last - 1 + cycleLen) % cycleLen // Исправляем возможный отрицательный индекс
		fmt.Println("index:", index, cycle[index])
		last = cycle[index]
	}

	return last
}

// powerMod возвращает x^y % mod
func powerMod(x, y int, mod int) int {
	result := 1
	base := x % mod

	for y > 0 {
		// Если степень нечетная, умножаем результат на base
		if y%2 == 1 {
			result = (result * base) % mod
		}
		// Увеличиваем степень
		y /= 2
		// Увеличиваем base
		base = (base * base) % mod
	}

	return result
}
