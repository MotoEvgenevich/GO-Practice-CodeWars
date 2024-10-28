package main

import "fmt"

func main() {
	fmt.Println(FindShort("Let's travel abroad shall we"))

}

func FindShort(s string) int {
	slice := []byte{}
	for _, value := range s {
		if value != 32 {
			slice = append(slice, byte(value))
		}
		if value == 32 {
			s := make([]byte, 3)
			fmt.Println(s)
		}
	}
	fmt.Println(string(slice))
	return 0
}
