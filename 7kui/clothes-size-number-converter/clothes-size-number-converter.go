package kata

func SizeToNumber(size string) (int, bool) {
	if len(size) == 0 {
		return 0, false
	}

	// Переменные для хранения базового размера и модификатора
	baseSize := 0
	multiplier := 0

	// Определяем базовый размер для строк, не содержащих модификаторов (например, "s", "m", "l")
	switch size[len(size)-1] {
	case 's':
		baseSize = 36
		multiplier = -2
	case 'm':
		// Если перед 'm' есть 'x', возвращаем невалидный результат
		if len(size) > 1 {
			return 0, false
		}
		return 38, true
	case 'l':
		baseSize = 40
		multiplier = 2
	default:
		return 0, false
	}

	// Проверяем наличие модификаторов 'x' в начале строки
	for i := 0; i < len(size)-1; i++ {
		if size[i] != 'x' {
			return 0, false
		}
	}

	// Если есть 'x' в строке, пересчитываем размер
	finalSize := baseSize + (multiplier * (len(size) - 1))

	// Возвращаем результат
	return finalSize, true
}

/*



Description

You have clothes international size as text (xs, s, xxl).
Your goal is to return European number size as a number from that size.

Notice that there is arbitrary amount of modifiers (x), excluding m size, and input can be any string.

Linearity

Base size for medium (m) is 38.
(then, small (s) is 36, and large (l) is 40)

The scale is linear; modifier x continues that by adding or subtracting 2 from resulting size.

(For sizes where x modifier makes total size negative, treat that as OK, and return negative size)

Invalid cases

Function should handle wrong/invalid sizes.

Valid input has any base size (s/m/l) and any amount of modifiers (x) before base size (like xxl).
Notice that you cannot apply modifier for m size, consider that case as invalid!
Anything else is disallowed and should be considered as invalid (None for Python, 0, false for Go, null for JavaScript).

Invalid examples: xxx, sss, xm, other string

Valid Examples

xs: 34
s: 36
m: 38
l: 40
xl: 42

*/
