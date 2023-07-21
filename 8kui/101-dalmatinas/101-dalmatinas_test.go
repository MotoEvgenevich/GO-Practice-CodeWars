package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Dogs     int
	Expected string
}

var testData = []SolveData{
	{
		Dogs:     26,
		Expected: "More than a handful!",
	},
	{
		Dogs:     8,
		Expected: "Hardly any",
	},
	{
		Dogs:     14,
		Expected: "More than a handful!",
	},
	{
		Dogs:     80,
		Expected: "Woah that's a lot of dogs!",
	},
	{
		Dogs:     100,
		Expected: "Woah that's a lot of dogs!",
	},
	{
		Dogs:     50,
		Expected: "More than a handful!",
	},
	{
		Dogs:     10,
		Expected: "Hardly any",
	},
	{
		Dogs:     101,
		Expected: "101 DALMATIONS!!!",
	},
}

func TestHowManyDalmatians(t *testing.T) {
	for _, test := range testData {
		output := HowManyDalmatians(test.Dogs)
		if len(output) != len(test.Expected) {
			t.Errorf("Expected length of output %v does not match expected length %v", len(output), len(test.Expected))
			continue
		}
		for i := 0; i < len(output); i++ {
			if output[i] != test.Expected[i] {
				t.Errorf("Output %v at index %d does not match expected %v", output[i], i, test.Expected[i])
			}
		}
	}
}

func BenchmarkHowManyDalmatians(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HowManyDalmatians(101)
	}
}

func ExampleHowManyDalmatians() {
	fmt.Println(HowManyDalmatians(101))
	// Output: 101 DALMATIONS!!!
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>
