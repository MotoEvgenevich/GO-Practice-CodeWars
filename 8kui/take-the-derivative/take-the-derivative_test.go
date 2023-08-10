package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Coefficient int
	Exponent    int
	Expected    string
}

var testData = []SolveData{
	{
		Coefficient: 7,
		Exponent:    8,
		Expected:    "56x^7",
	},
	{
		Coefficient: 5,
		Exponent:    9,
		Expected:    "45x^8",
	},
}

func TestHowManyDerive(t *testing.T) {
	for _, test := range testData {
		output := Derive(test.Coefficient, test.Exponent)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
		}
	}
}

func BenchmarkDerive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Derive(7, 8)
	}
}

func ExampleDerive() {
	fmt.Println(Derive(7, 8))
	// Output: 56x^7
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
derive(7, 8) --> this should output "56x^7"
derive(5, 9) --> this should output "45x^8" */
