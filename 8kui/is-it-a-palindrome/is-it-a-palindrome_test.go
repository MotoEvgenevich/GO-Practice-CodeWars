package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Str      string
	Expected bool
}

var testData = []SolveData{
	{
		Str:      "a",
		Expected: true,
	},
	{
		Str:      "aba",
		Expected: true,
	},
	{
		Str:      "Abba",
		Expected: true,
	},
	{
		Str:      "hello",
		Expected: false,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range testData {
		output := IsPalindrome(test.Str)
		if output != test.Expected {
			t.Errorf("Value of output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("hello")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("hello"))
	// Output: false
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*



Expect(IsPalindrome("a")).To(Equal(true))
Expect(IsPalindrome("aba")).To(Equal(true))
Expect(IsPalindrome("Abba")).To(Equal(true))
Expect(IsPalindrome("hello")).To(Equal(false)) */
