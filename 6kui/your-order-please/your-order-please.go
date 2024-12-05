package kata

import (
	"strconv"
	"strings"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")
	result := []string{}
	for i := 0; i <= len(words); i++ {
		for _, v := range words {
			if strings.Contains(v, strconv.Itoa(i)) {
				result = append(result, v)
			}
		}
	}
	return strings.Join(result, " ")
}
