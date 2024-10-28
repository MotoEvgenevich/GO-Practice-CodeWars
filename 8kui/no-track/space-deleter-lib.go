package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NoSpace("word word word"))
}

func NoSpace(word string) string {
	result := strings.ReplaceAll(word, " ", "")
	return result
}

// первая дополнительная задача —
// вспомним что ты пару дней назад сам написал
// функцию ReplaceAll и используем ее тут

// и потом также используем strings.ReplaceAll
// из стандартной либы — надо будет найти тут
// скриншот что ты скидывал из чужого решения
// и понять почему у них работало, а у тебя — нет
