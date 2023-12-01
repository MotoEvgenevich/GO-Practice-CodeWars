package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	N        int
	Expected int
}

var testData = []SolveData{
	{
		N:        999,
		Expected: 4,
	},
	{
		N:        39,
		Expected: 3,
	},
	{
		N:        4,
		Expected: 0,
	},
	{
		N:        25,
		Expected: 2,
	},
	{
		N:        0,
		Expected: 0,
	},
}

func TestHowManyDalmatians(t *testing.T) {
	for _, test := range testData {
		output := Persistence(test.N)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkPersistence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Persistence(999)
	}
}

func ExamplePersistence() {
	fmt.Println(Persistence(999))
	// Output: 4
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
It("Case 1", func() {
    Expect(Persistence(39)).To(Equal(3))
  })
  It("Case 2", func() {
    Expect(Persistence(4)).To(Equal(0))
  })
  It("Case 3", func() {
    Expect(Persistence(25)).To(Equal(2))
  })
  It("Case 4", func() {
    Expect(Persistence(999)).To(Equal(4))
  })
*/
