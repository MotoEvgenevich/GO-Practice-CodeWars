package kata

func Parse(data string) []int {
	digit := 0
	arr := []int{}
	for _, v := range data {
		switch v {
		case 'i':
			digit++
		case 'd':
			digit--
		case 's':
			digit *= digit
		case 'o':
			arr = append(arr, digit)
		}
	}
	return arr
}

/*
Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/
