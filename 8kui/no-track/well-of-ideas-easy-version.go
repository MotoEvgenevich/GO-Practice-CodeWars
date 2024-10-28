package main

import "fmt"

func main() {
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}))
}

func Well(x []string) string {
	goodSlice := []string{}
	for _, cond := range x {
		if cond == "good" {
			goodSlice = append(goodSlice, cond)

		}
		fmt.Println(len(goodSlice))
	}
	if len(goodSlice) == 1 || len(goodSlice) == 2 {
		return "Publish!"
	}
	if len(goodSlice) > 2 {
		return "I smell a series!"
	} else {
		return "Fail!"
	}

}

func WellCase(x []string) string {
	good := 0
	for _, idea := range x {
		switch idea {
		case "good":
			good++
		}
	}
	switch {
	case good > 2:
		return "I smell a series!"
	case good > 0:
		return "Publish!"
	default:
		return "Fail!"
	}
}
