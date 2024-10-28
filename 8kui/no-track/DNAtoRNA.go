package main

import "fmt"

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}

func DNAtoRNA(dna string) string {
	// your code here
	outputSlice := []byte{}
	T := "T"
	for i := 0; i < len(dna); i++ {
		if string(dna[i]) == T {
			outputSlice = append(outputSlice, "U"...)
		} else {
			outputSlice = append(outputSlice, dna[i])
		}
	}

	return string(outputSlice)
}

// "GCAT" => "GCAU"
