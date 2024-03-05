package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Myl      [][2]int
	Interval [2]int
	Expected [][2]int
}

var testData = []SolveData{
	{
		Myl:      [][2]int{{1, 2}},
		Interval: [2]int{3, 4},
		Expected: [][2]int{{1, 2}, {3, 4}},
	},
	{

		Myl:      [][2]int{{1, 2}, {3, 4}, {5, 6}},
		Interval: [2]int{2, 3},
		Expected: [][2]int{{1, 4}, {5, 6}},
	},
}

func TestIntervalInsert(t *testing.T) {
	for _, test := range testData {
		output := IntervalInsert(test.Myl, test.Interval)
		if !compareSlices(output, test.Expected) {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
		}
	}
}

func compareSlices(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkIntervalInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntervalInsert([][2]int{{1, 2}, {3, 4}, {5, 6}}, [2]int{2, 3})
	}
}

func ExampleIntervalInsert() {
	fmt.Println(IntervalInsert([][2]int{{1, 2}, {3, 4}, {5, 6}}, [2]int{2, 3}))
	// Output: [[1 4] [5 6]]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* It("Testing [[1,2]] and [3,4]",func() {Expect(IntervalInsert([][2]int{{1,2}},[2]int{3,4})).To(Equal([][2]int{{1,2},{3,4}}))})
   It("Testing [[1,2], [3,4]] and [2,3]",func() {Expect(IntervalInsert([][2]int{{1,2},{3,4}},[2]int{2,3})).To(Equal([][2]int{{1,4}}))})
   It("Testing [[1,2], [3,4], [5,6]] and [2,3]",func() {Expect(IntervalInsert([][2]int{{1,2},{3,4},{5,6}},[2]int{2,3})).To(Equal([][2]int{{1,4},{5,6}}))})
   Test inputs: [[188 191] [198 199] [208 211] [218 219] [228 231] [238 239] [248 250] [258 261] [268 269] [278 280] [288 289] [298 299] [308 310] [318 320] [328 331] [338 341] [348 349]] and [216 293]
*/
