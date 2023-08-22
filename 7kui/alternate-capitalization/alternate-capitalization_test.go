package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	St       string
	Expected []string
}

var tests = []SolveTest{
	{
		St:       "abcdef",
		Expected: []string{"AbCdEf", "aBcDeF"},
	},
	{
		St:       "codewars",
		Expected: []string{"CoDeWaRs", "cOdEwArS"},
	},
	{
		St:       "abracadabra",
		Expected: []string{"AbRaCaDaBrA", "aBrAcAdAbRa"},
	},
	{
		St:       "codewarriors",
		Expected: []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"},
	},
	{
		St:       "indexinglessons",
		Expected: []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"},
	},
	{
		St:       "codingisafunactivity",
		Expected: []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"},
	},
}

func TestCapitalize(t *testing.T) {
	for _, test := range tests {
		output := Capitalize(test.St)
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

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Capitalize("codingisafunactivity")
	}
}

func ExampleCapitalize() {
	fmt.Println(Capitalize("codewars"))
	// Output: [CoDeWaRs cOdEwArS]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* var _ = Describe("Example tests", func() {
	It("It should work for basic tests", func() {
		  dotest("abcdef", []string{"AbCdEf", "aBcDeF"})
		  dotest("codewars", []string{"CoDeWaRs", "cOdEwArS"})
		  dotest("abracadabra", []string{"AbRaCaDaBrA", "aBrAcAdAbRa"})
		  dotest("codewarriors", []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"})
		  dotest("indexinglessons", []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"})
		  dotest("codingisafunactivity", []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"})
	})
  }) */
