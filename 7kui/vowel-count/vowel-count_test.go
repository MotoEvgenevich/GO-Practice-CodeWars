package getvowelcount

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Str      string
	Expected int
}

var testData = []SolveData{
	{
		Str:      "should test that the solution returns the correct value	",
		Expected: 17,
	},
	{
		Str:      "mcnwnot zonmly pjqikoig xwukdfj ft tthgie ue...",
		Expected: 10,
	},
	{
		Str:      "abracadabra",
		Expected: 5,
	},
	{
		Str:      "pmyhucyn lgbfkjn cpdln qodnhi hfwldqbq jdtnmtd kgw...",
		Expected: 3,
	},
	{
		Str:      "xcgfni q bzjcblj rxrfhql psufswl qju rphmv...",
		Expected: 3,
	},
	{
		Str:      "cyckvsrc cr ejdmugf yqbww iptjiwws sgxq	",
		Expected: 4,
	},
}

func TestToGetCount(t *testing.T) {
	for _, test := range testData {
		output := GetCount(test.Str)
		if output != test.Expected {
			t.Errorf("For input %v, output %v does not match expected %v", test.Str, output, test.Expected)
		}
	}
}

func BenchmarkGetCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCount("pmyhucyn lgbfkjn cpdln qodnhi hfwldqbq jdtnmtd kgw...")
	}
}

func ExampleGetCount() {
	fmt.Println(GetCount("should test that the solution returns the correct value"))
	// Output: 17
}

func CoverMode() string {
	mode := testing.CoverMode()
	if mode == "" {
		return "Тестовое покрытие не активировано."
	}
	return fmt.Sprintf("Режим тестового покрытия установлен на: %s", mode)
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
should test that the solution returns the correct value			17
mcnwnot zonmly pjqikoig xwukdfj ft tthgie ue...					10
abracadabra														5
pmyhucyn lgbfkjn cpdln qodnhi hfwldqbq jdtnmtd kgw...			3
xcgfni q bzjcblj rxrfhql psufswl qju rphmv...					3
cyckvsrc cr ejdmugf yqbww iptjiwws sgxq							4
*/
