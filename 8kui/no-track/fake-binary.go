package main

import "fmt"

func main() {
	fmt.Println(FakeBin("366058562030849490134388085"))
}

func FakeBin(x string) string {
	slice := []byte{}
	bin := ""
	for _, v := range x {
		slice = append(slice, byte(v))
	}
	fmt.Println(slice)
	for _, v := range slice {
		switch {
		case v < 53:
			bin += "0"
		case v >= 53:
			bin += "1"

		}
	}
	return bin
}
