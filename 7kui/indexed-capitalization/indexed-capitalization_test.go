package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Str      string
	Arr      []int
	Expected string
}

var testData = []SolveTest{
	{
		Str:      "abcdef",
		Arr:      []int{1, 2, 5},
		Expected: "aBCdeF",
	},
	{
		Str:      "abcdef",
		Arr:      []int{1, 2, 5, 100},
		Expected: "aBCdeF",
	},
	{
		Str:      "codewars",
		Arr:      []int{1, 3, 5, 50},
		Expected: "cOdEwArs",
	},
	{
		Str:      "abracadabra",
		Arr:      []int{2, 6, 9, 10},
		Expected: "abRacaDabRA",
	},
	{
		Str:      "codewarriors",
		Arr:      []int{5},
		Expected: "codewArriors",
	},
	{
		Str:      "indexinglessons",
		Arr:      []int{0},
		Expected: "Indexinglessons",
	},
}

func TestTCapitalize(t *testing.T) {
	for _, test := range testData {
		output := Capitalize(test.Str, test.Arr)
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

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Capitalize("indexinglessons", []int{0})
	}
}

func ExampleCapitalize() {
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5}))
	// Output: aBCdeF
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* var _ = Describe("Example tests", func() {
It("It should work for basic tests", func() {
	  dotest("abcdef", []int {1,2,5},"aBCdeF" )
	  dotest( "abcdef",[]int {1,2,5,100},"aBCdeF")
	  dotest("codewars", []int {1,3,5,50},"cOdEwArs" )
	  dotest("abracadabra", []int {2,6,9,10},"abRacaDabRA")
	  dotest("codewarriors",[]int {5},"codewArriors")
	  dotest( "indexinglessons",[]int {0},"Indexinglessons")
}) */
