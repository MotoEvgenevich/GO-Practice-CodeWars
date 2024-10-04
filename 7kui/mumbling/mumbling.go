package kata

import (
	"strings"
)

func Accum(s string) string {
	s = strings.ToUpper(s)
	type letters struct {
		odrer  int
		symbol string
	}
	someSting := ""
	for i, v := range s {
		p1 := letters{
			odrer:  i,
			symbol: string(v),
		}
		someSting += p1.symbol
		for i := 0; i < p1.odrer; i++ {
			someSting += strings.ToLower(p1.symbol)
		}
		someSting += "-"
	}
	someSting = someSting[:(len(someSting) - 1)]
	return someSting
}

/*
This time no story, no theory. The examples below show you how to write function accum:

Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/
