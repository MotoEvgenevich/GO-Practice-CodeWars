package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Arr      []int
	Expected []int
}

var testData = []SolveData{
	{
		Arr:      []int{2, 3, 5, 3, 7, 9, 5, 3, 7},
		Expected: []int{3, 3, 3, 5, 5, 7, 7, 2, 9},
	},
	{
		Arr:      []int{1, 2, 3, 0, 5, 0, 1, 6, 8, 8, 6, 9, 1},
		Expected: []int{1, 1, 1, 0, 0, 6, 6, 8, 8, 2, 3, 5, 9},
	},
	{
		Arr:      []int{5, 9, 6, 9, 6, 5, 9, 9, 4, 4},
		Expected: []int{9, 9, 9, 9, 4, 4, 5, 5, 6, 6},
	},
	{
		Arr:      []int{4, 4, 2, 5, 1, 1, 3, 3, 2, 8},
		Expected: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 8},
	},
	{
		Arr:      []int{4, 9, 5, 0, 7, 3, 8, 4, 9, 0},
		Expected: []int{0, 0, 4, 4, 9, 9, 3, 5, 7, 8},
	},
}

func TestHowManyDalmatians(t *testing.T) {
	for _, test := range testData {
		output := Solve(test.Arr)
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

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7})
	}
}

func ExampleSolve() {
	fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7}))
	// Output: [3 3 3 5 5 7 7 2 9]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* var _ = Describe("Sample tests", func() {
	It("should handle basic cases", func() {
		Expect(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7})).To(Equal([]int{3, 3, 3, 5, 5, 7, 7, 2, 9}))
		Expect(Solve([]int{1, 2, 3, 0, 5, 0, 1, 6, 8, 8, 6, 9, 1})).To(Equal([]int{1, 1, 1, 0, 0, 6, 6, 8, 8, 2, 3, 5, 9}))
		Expect(Solve([]int{5, 9, 6, 9, 6, 5, 9, 9, 4, 4})).To(Equal([]int{9, 9, 9, 9, 4, 4, 5, 5, 6, 6}))
		Expect(Solve([]int{4, 4, 2, 5, 1, 1, 3, 3, 2, 8})).To(Equal([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 8}))
		Expect(Solve([]int{4, 9, 5, 0, 7, 3, 8, 4, 9, 0})).To(Equal([]int{0, 0, 4, 4, 9, 9, 3, 5, 7, 8}))
	})
})
*/
