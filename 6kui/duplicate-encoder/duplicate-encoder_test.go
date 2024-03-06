package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Word     string
	Expected string
}

var testData = []SolveData{
	{
		Word:     "din",
		Expected: "(((",
	},
	{
		Word:     "recede",
		Expected: "()()()",
	},
	{
		Word:     "(( @",
		Expected: "))((",
	},
	{
		Word:     "Success",
		Expected: ")())())",
	},
}

func TestDuplicateEncode(t *testing.T) {
	for _, test := range testData {
		output := DuplicateEncode(test.Word)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkDuplicateEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DuplicateEncode("Success")
	}
}

func ExampleDuplicateEncode() {
	fmt.Println(DuplicateEncode("Success"))
	// Output: )())())
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
TEST CASES
    Expect(DuplicateEncode("din")).To(Equal("((("))
    Expect(DuplicateEncode("recede")).To(Equal("()()()"))
    Expect(DuplicateEncode("(( @")).To(Equal("))(("))
  })

  It("should ignore case", func() {
    Expect(DuplicateEncode("Success")).To(Equal(")())())"))
*/
