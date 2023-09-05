package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Arr      []string
	Expected string
}

var testData = []SolveData{
	{
		Arr:      []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"},
		Expected: "b***i***t***c***o***i***n",
	},
	{
		Arr:      []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"},
		Expected: "a***r***e",
	},
	{
		Arr:      []string{"lets", "talk", "about", "javascript", "the", "best", "language"},
		Expected: "a***b***o***u***t",
	},
	{
		Arr:      []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"},
		Expected: "c***o***d***e",
	},
	{
		Arr:      []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"},
		Expected: "L***e***t***s",
	},
}

func TestTwoSort(t *testing.T) {
	for _, test := range testData {
		output := TwoSort(test.Arr)
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

func BenchmarkTwoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSort([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"})
	}
}

func ExampleTwoSort() {
	fmt.Println(TwoSort([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}))
	// Output: a***r***e
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*

TEST CASES:

s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
Expect(TwoSort(s)).To(Equal("a***r***e"))
s = []string{"lets", "talk", "about", "javascript", "the", "best", "language"}
Expect(TwoSort(s)).To(Equal("a***b***o***u***t"))
s = []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}
Expect(TwoSort(s)).To(Equal("c***o***d***e"))
s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
Expect(TwoSort(s)).To(Equal("L***e***t***s")) */
