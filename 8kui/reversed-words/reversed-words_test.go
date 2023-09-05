package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Str      string
	Expected string
}

var testData = []SolveData{
	{
		Str:      "yoda doesn't speak like this",
		Expected: "this like speak doesn't yoda",
	},
	{
		Str:      "hello world!",
		Expected: "world! hello",
	},
	{
		Str:      "foobar",
		Expected: "foobar",
	},
	{
		Str:      "kata editor",
		Expected: "editor kata",
	},
	{
		Str:      "row row row your boat",
		Expected: "boat your row row row",
	},
}

func TestReverseWords(t *testing.T) {
	for _, test := range testData {
		output := ReverseWords(test.Str)
		if len(output) != len(test.Expected) {
			t.Errorf("Expected length of output %v does not match expected length %v", len(output), len(test.Expected))

		}
		for i, v := range output {
			if v != rune(test.Expected[i]) {
				t.Errorf("Output %v does not match Expected %v", output, test.Expected)
			}
		}
	}
}

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseWords("hello world!")
	}
}

func ExampleReverseWords() {
	fmt.Println(ReverseWords("hello world!"))
	// Output: world! hello
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* Expect(ReverseWords("hello world!")).To(Equal("world! hello"))
Expect(ReverseWords("yoda doesn't speak like this")).To(Equal("this like speak doesn't yoda"))
Expect(ReverseWords("foobar")).To(Equal("foobar"))
Expect(ReverseWords("kata editor")).To(Equal("editor kata"))
Expect(ReverseWords("row row row your boat")).To(Equal("boat your row row row")) */
