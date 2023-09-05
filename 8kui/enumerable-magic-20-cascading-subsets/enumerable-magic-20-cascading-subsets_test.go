package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Arr      []int
	n        int
	Expected [][]int
}

var testData = []SolveData{
	{
		Arr:      []int{3, 5, 8, 13},
		n:        1,
		Expected: [][]int{{3}, {5}, {8}, {13}},
	},
	{
		Arr:      []int{3, 5, 8, 13},
		n:        2,
		Expected: [][]int{{3, 5}, {5, 8}, {8, 13}},
	},
	{
		Arr:      []int{3, 5, 8, 13},
		n:        3,
		Expected: [][]int{{3, 5, 8}, {5, 8, 13}},
	},
}

func TestEachCons(t *testing.T) {
	for _, test := range testData {
		output := EachCons(test.Arr, test.n)
		if len(output) != len(test.Expected) {
			t.Errorf("Expected value of output %v does not match test value %v", len(output), len(test.Expected))

		}
	}
}

func BenchmarkEachCons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EachCons([]int{3, 5, 8, 13}, 3)
	}
}

func ExampleEachCons() {
	fmt.Println(EachCons([]int{3, 5, 8, 13}, 3))
	// Output: [[3 5 8] [5 8 13]]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
 TEST CASES:

 arr := []int{3, 5, 8, 13}
       Expect(EachCons(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
       Expect(EachCons(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
       Expect(EachCons(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
       Expect(EachCons([]int{}, 3)).To(BeEmpty())
*/
