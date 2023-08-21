package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	array    [][]int
	Expected string
}

var testData = []SolveData{
	{
		array:    [][]int{{0, 1, 2, 3, 45}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}},
		Expected: "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34",
	},
	{
		array:    [][]int{{-25, 21, 2, -33, 48}, {30, 31, -32, 33, -34}},
		Expected: "-25,21,2,-33,48\n30,31,-32,33,-34",
	},
	{
		array:    [][]int{{5, 55, 5, 5, 55}, {6, 6, 66, 23, 24}, {666, 31, 66, 33, 7}},
		Expected: "5,55,5,5,55\n6,6,66,23,24\n666,31,66,33,7",
	},
}

func TestToCsvText(t *testing.T) {
	for _, test := range testData {
		output := ToCsvText(test.array)
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

func BenchmarkToCsvText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToCsvText([][]int{{0, 1, 2, 3, 45}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}})
	}
}

func ExampleToCsvText() {
	fmt.Println(ToCsvText([][]int{{0, 1, 2, 3, 45}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}}))
	// Output: 0,1,2,3,45
	//10,11,12,13,14
	//20,21,22,23,24
	//30,31,32,33,34
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* Expect(ToCsvText([][]int{
	{0, 1, 2, 3, 45},
	{10, 11, 12, 13, 14},
	{20, 21, 22, 23, 24},
	{30, 31, 32, 33, 34}})).To(Equal("0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34"))
Expect(ToCsvText([][]int{
	{-25, 21, 2, -33, 48},
	{30, 31, -32, 33, -34}})).To(Equal("-25,21,2,-33,48\n30,31,-32,33,-34"))
Expect(ToCsvText([][]int{
	{5, 55, 5, 5, 55},
	{6, 6, 66, 23, 24},
	{666, 31, 66, 33, 7}})).To(Equal("5,55,5,5,55\n6,6,66,23,24\n666,31,66,33,7")) */
