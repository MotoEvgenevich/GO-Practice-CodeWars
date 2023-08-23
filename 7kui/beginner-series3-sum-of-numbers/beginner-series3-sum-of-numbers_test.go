package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	a        int
	b        int
	Expected int
}

var tests = []SolveTest{
	{
		a:        0,
		b:        1,
		Expected: 1,
	},
	{
		a:        1,
		b:        2,
		Expected: 3,
	},
	{
		a:        5,
		b:        -1,
		Expected: 14,
	},
	{
		a:        505,
		b:        4,
		Expected: 127759,
	},
	{
		a:        321,
		b:        123,
		Expected: 44178,
	},
	{
		a:        -50,
		b:        0,
		Expected: -1275,
	},
	{
		a:        -1,
		b:        -5,
		Expected: -15,
	},
	{
		a:        -5,
		b:        -5,
		Expected: -5,
	},
	{
		a:        -505,
		b:        4,
		Expected: -127755,
	},
	{
		a:        -321,
		b:        123,
		Expected: -44055,
	},
	{
		a:        0,
		b:        0,
		Expected: 0,
	},
	{
		a:        -5,
		b:        -1,
		Expected: -15,
	},
	{
		a:        0,
		b:        1,
		Expected: 1,
	},
	{
		a:        0,
		b:        1,
		Expected: 1,
	},
	{
		a:        0,
		b:        1,
		Expected: 1,
	},
}

func TestPrinterError(t *testing.T) {
	for _, test := range tests {
		output := GetSum(test.a, test.b)
		if output != test.Expected {
			t.Errorf("Expected %v does not match with output %v", output, test.Expected)
		}
	}
}

func BenchmarkGetSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSum(-505, 4)
	}
}

func ExampleGetSum() {
	fmt.Println(GetSum(-505, 4))
	// Output: -127755
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
Expect(GetSum(0, 1)).To(Equal(1))
Expect(GetSum(1, 2)).To(Equal(3))
Expect(GetSum(5, -1)).To(Equal(14))
Expect(GetSum(505, 4)).To(Equal(127759))
Expect(GetSum(321, 123)).To(Equal(44178))
Expect(GetSum(0, -1)).To(Equal(-1))
Expect(GetSum(-50, 0)).To(Equal(-1275))
Expect(GetSum(-1, -5)).To(Equal(-15))
Expect(GetSum(-5, -5)).To(Equal(-5))
Expect(GetSum(-505, 4)).To(Equal(-127755))
Expect(GetSum(-321, 123)).To(Equal(-44055))
Expect(GetSum(0, 0)).To(Equal(0))
Expect(GetSum(-5, -1)).To(Equal(-15))
Expect(GetSum(5, 1)).To(Equal(15))
Expect(GetSum(-17, -17)).To(Equal(-17))
Expect(GetSum(17, 17)).To(Equal(17)) */
