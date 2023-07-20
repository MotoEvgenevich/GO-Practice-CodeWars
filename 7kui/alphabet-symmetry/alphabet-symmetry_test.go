package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	MySlice  []string
	Expected []int
}

var tests = []SolveTest{
	{
		MySlice:  []string{"abode", "ABc", "xyzD"},
		Expected: []int{4, 3, 1},
	},
	{
		MySlice:  []string{"abide", "ABc", "xyz"},
		Expected: []int{4, 3, 0},
	},
	{
		MySlice:  []string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"},
		Expected: []int{6, 5, 7},
	},
	{
		MySlice:  []string{"encode", "abc", "xyzD", "ABmD"},
		Expected: []int{1, 3, 1, 3},
	},
}

func TestSolve(t *testing.T) {
	for _, test := range tests {
		output := Solve(test.MySlice)
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
		Solve([]string{"encode", "abc", "xyzD", "ABmD"})
	}
}

func ExampleSolve() {
	fmt.Println(Solve([]string{"encode", "abc", "xyzD", "ABmD"}))
	// Output: [1 3 1 3]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* 	Expect(solve([]string{"abode","ABc","xyzD"})).To(Equal([]int{4,3,1}))
Expect(solve([]string{"abide","ABc","xyz"})).To(Equal([]int{4,3,0}))
Expect(solve([]string{"IAMDEFANDJKL","thedefgh","xyzDEFghijabc"})).To(Equal([]int{6,5,7}))
Expect(solve([]string{"encode","abc","xyzD","ABmD"})).To(Equal([]int{1,3,1,3})) */
