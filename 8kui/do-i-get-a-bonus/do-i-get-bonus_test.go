package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Salary   int
	Bonus    bool
	Expected string
}

var testData = []SolveData{
	{
		Salary:   26,
		Bonus:    true,
		Expected: "£260",
	},
	{
		Salary:   100,
		Bonus:    true,
		Expected: "£1000",
	},
	{
		Salary:   14000,
		Bonus:    true,
		Expected: "£140000",
	},
	{
		Salary:   80,
		Bonus:    false,
		Expected: "£80",
	},
}

func TestHowManyDalmatians(t *testing.T) {
	for _, test := range testData {
		output := BonusTime(test.Salary, test.Bonus)
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

func BenchmarkHowManyDalmatians(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BonusTime(101, true)
	}
}

func ExampleBonusTime() {
	fmt.Println(BonusTime(101, false))
	// Output: £101
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>
