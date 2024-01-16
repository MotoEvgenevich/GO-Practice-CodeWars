package kata

import (
	"strings"
)

func Solve(s string, a, b int) string {
	var builder strings.Builder
	// Проверка и корректировка индексов
	if b >= len(s) {
		b = len(s) - 1
	}
	if a >= len(s) {
		a = len(s) - 1
	}

	// Добавление начальной части строки
	for i := 0; i < a; i++ {
		builder.WriteString(string(s[i]))
	}

	// Реверсирование средней части
	for i := b; i >= a; i-- {
		builder.WriteString(string(s[i]))
	}

	// Добавление оставшейся части строки
	if b < len(s) {
		for i := b + 1; i < len(s); i++ {
			builder.WriteString(string(s[i]))
		}
	}
	return builder.String()
}

/*
In this Kata, you will be given a string and two indexes (a and b).
Your task is to reverse the portion of that string between those two
indices inclusive.

solve("codewars",1,5) = "cawedors" -- elements at index 1 to 5 inclusive are "odewa". So we reverse them.
solve("cODEWArs", 1,5) = "cAWEDOrs" -- to help visualize.
Input will be lowercase and uppercase letters only.

The first index a will always be lower that than the string length; the second index b can be greater than the string length. More examples in the test cases. Good luck! */
