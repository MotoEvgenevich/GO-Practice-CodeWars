package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Number   int
	Expected string
}

var testData = []SolveData{
	{
		Number:   5,
		Expected: "1 * 5 = 5\n2 * 5 = 10\n3 * 5 = 15\n4 * 5 = 20\n5 * 5 = 25\n6 * 5 = 30\n7 * 5 = 35\n8 * 5 = 40\n9 * 5 = 45\n10 * 5 = 50",
	},
	{
		Number:   1,
		Expected: "1 * 1 = 1\n2 * 1 = 2\n3 * 1 = 3\n4 * 1 = 4\n5 * 1 = 5\n6 * 1 = 6\n7 * 1 = 7\n8 * 1 = 8\n9 * 1 = 9\n10 * 1 = 10",
	},
	{
		Number:   2,
		Expected: "1 * 2 = 2\n2 * 2 = 4\n3 * 2 = 6\n4 * 2 = 8\n5 * 2 = 10\n6 * 2 = 12\n7 * 2 = 14\n8 * 2 = 16\n9 * 2 = 18\n10 * 2 = 20",
	},
}

func TestMultiTable(t *testing.T) {
	for _, test := range testData {
		output := MultiTable(test.Number)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
		}
	}
}

func BenchmarkMultiTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiTable(9)
	}
}

func ExampleMultiTable() {
	fmt.Println(MultiTable(5))
	// Output: 1 * 5 = 5
	//2 * 5 = 10
	//3 * 5 = 15
	//4 * 5 = 20
	//5 * 5 = 25
	//6 * 5 = 30
	//7 * 5 = 35
	//8 * 5 = 40
	//9 * 5 = 45
	//10 * 5 = 50
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* Expect(MultiTable(5)).To(Equal("1 * 5 = 5\n2 * 5 = 10\n3 * 5 = 15\n4 * 5 = 20\n5 * 5 = 25\n6 * 5 = 30\n7 * 5 = 35\n8 * 5 = 40\n9 * 5 = 45\n10 * 5 = 50"))
Expect(MultiTable(1)).To(Equal("1 * 1 = 1\n2 * 1 = 2\n3 * 1 = 3\n4 * 1 = 4\n5 * 1 = 5\n6 * 1 = 6\n7 * 1 = 7\n8 * 1 = 8\n9 * 1 = 9\n10 * 1 = 10")) */
