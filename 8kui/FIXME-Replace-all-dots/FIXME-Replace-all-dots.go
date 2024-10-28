package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(ReplaceDots("one.two.three"))
}

func ReplaceDots(str string) string {
	fmt.Println(len(str))
	return regexp.MustCompile(`\.`).ReplaceAllString(str, "-")
}

/* The code provided is supposed replace all the dots . in the specified String str with dashes -

But it's not working properly.

Task

Fix the bug so we can all go home early.

Notes

String str will never be null. */
