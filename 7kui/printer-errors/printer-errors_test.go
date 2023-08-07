package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	ColorString string
	Expected    string
}

var tests = []SolveTest{
	{
		ColorString: "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz",
		Expected:    "3/56",
	},
	{
		ColorString: "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz",
		Expected:    "6/60",
	},
	{
		ColorString: "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu",
		Expected:    "11/65",
	},
}

func TestPrinterError(t *testing.T) {
	for _, test := range tests {
		output := PrinterError(test.ColorString)
		if output != test.Expected {
			t.Errorf("Expected %v does not match with output %v", output, test.Expected)
		}
	}
}

func BenchmarkPrinterError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")
	}
}

func ExamplePrinterError() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	// Output: 3/56
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* 	Expect(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("3/56"))
Expect(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("6/60"))
Expect(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")).To(Equal("11/65")) */
