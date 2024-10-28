/*
All of the animals are having a feast! Each animal is bringing one dish.
There is just one rule: the dish must start and end with the same letters as the animal's name.
For example, the great blue heron is bringing garlic naan and the chickadee is bringing chocolate cake.

Write a function feast that takes the animal's name and dish as arguments
and returns true or false to indicate whether the beast is allowed to bring the dish to the feast.

Assume that beast and dish are always lowercase strings,
and that each has at least two letters. beast and dish may contain hyphens and spaces,
but these will not appear at the beginning or end of the string. They will not contain numerals.

INPUT DATA:
"chickadee", "chocolate cake"
"great blue heron", "garlic naan"
"brown bear", "bear claw"
*/

package main

import "fmt"

func main() {
	fmt.Println(Feast("brown bear", "bear claw"))
}

func Feast(beast string, dish string) bool {
	beastSlice := []byte{}
	for _, v := range beast {
		beastSlice = append(beastSlice, byte(v))
	}
	fmt.Println(beastSlice[0])
	fmt.Println(beastSlice[len(beastSlice)-1])

	dishSlice := []byte{}
	for _, v := range dish {
		dishSlice = append(dishSlice, byte(v))
	}

	if beastSlice[0] == dishSlice[0] && beastSlice[len(beastSlice)-1] == dishSlice[len(dishSlice)-1] {
		return true
	} else {
		return false
	}

}
