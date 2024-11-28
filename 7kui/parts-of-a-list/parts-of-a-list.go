package kata

import "strings"

func PartList(arr []string) string {
	var result []string
	for i := 1; i < len(arr); i++ {
		part1 := strings.Join(arr[:i], " ")
		part2 := strings.Join(arr[i:], " ")
		result = append(result, "("+part1+", "+part2+")")
	}
	return strings.Join(result, "")
}

/*
Write a function partlist that gives all the ways to divide a list (an array)
of at least two elements into two non-empty parts.

Each two non empty parts will be in a pair (or an array for languages without tuples or
a structin C - C: see Examples test Cases - )
Each part will be in a string
Elements of a pair must be in the same order as in the original array.

(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")
*/
