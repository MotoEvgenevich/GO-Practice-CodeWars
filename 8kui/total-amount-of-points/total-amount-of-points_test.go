package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	MySlice  []string
	Expected int
}

var tests = []SolveTest{
	{
		MySlice:  []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"},
		Expected: 30,
	},
	{
		MySlice:  []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"},
		Expected: 10,
	},
	{
		MySlice:  []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"},
		Expected: 0,
	},
	{
		MySlice:  []string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"},
		Expected: 15,
	},
	{
		MySlice:  []string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"},
		Expected: 12,
	},
}

func TestPoints(t *testing.T) {
	for _, test := range tests {
		output := Points(test.MySlice)
		if output != test.Expected {
			t.Errorf("resutl of output doesnt match Expected %v , %v", output, test.Expected)
			return
		}
	}
}

func BenchmarkPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	}
}

func ExamplePoints() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
	// Output: 30
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
     Expect(Points([]string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"})).To(Equal(30))
     Expect(Points([]string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"})).To(Equal(10))
     Expect(Points([]string{"0:1","0:2","0:3","0:4","1:2","1:3","1:4","2:3","2:4","3:4"})).To(Equal(0))
     Expect(Points([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})).To(Equal(15))
     Expect(Points([]string{"1:0","2:0","3:0","4:4","2:2","3:3","1:4","2:3","2:4","3:4"})).To(Equal(12))
})*/
