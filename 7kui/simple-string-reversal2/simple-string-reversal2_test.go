package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	S        string
	a        int
	b        int
	Expected string
}

var tests = []SolveTest{
	{
		S:        "codewars",
		a:        1,
		b:        5,
		Expected: "cawedors",
	},
	{
		S:        "codingIsFun",
		a:        2,
		b:        100,
		Expected: "conuFsIgnid",
	},
	{
		S:        "FunctionalProgramming",
		a:        2,
		b:        15,
		Expected: "FuargorPlanoitcnmming",
	},
	{
		S:        "FunctionalProgramming",
		a:        2,
		b:        15,
		Expected: "FuargorPlanoitcnmming",
	},
	{
		S:        "FunctionalProgramming",
		a:        2,
		b:        15,
		Expected: "FuargorPlanoitcnmming",
	},
}

func TestRoundSolve(t *testing.T) {
	for _, test := range tests {
		output := Solve(test.S, test.a, test.b)
		if output != test.Expected {
			t.Errorf("Expected of output %v does not match expected value %v", output, test.Expected)
			continue
		}
		for i := 0; i < len(output); i++ {
			if output[i] != test.Expected[i] {
				t.Errorf("Output %v at index %d does not match expected %v", output[i], i, test.Expected[i])
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("abcefghijklmnopqrstuvwxyz", 5, 20)
	}
}

func ExampleSolve() {
	fmt.Println(Solve("codewars", 1, 5))
	// Output: cawedors
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
dotest("codewars",1,5,"cawedors")
dotest("codingIsFun",2,100, "conuFsIgnid")
dotest("FunctionalProgramming", 2, 15,"FuargorPlanoitcnmming")
dotest("abcefghijklmnopqrstuvwxyz",0,20, "vutsrqponmlkjihgfecbawxyz")
dotest("abcefghijklmnopqrstuvwxyz",5,20, "abcefvutsrqponmlkjihgwxyz")  */
