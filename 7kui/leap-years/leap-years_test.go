package kata

import "testing"

func TestIsLeapYear(t *testing.T) {

	{
		got := IsLeapYear(2016)
		want := true
		if got != want {
			t.Errorf("IsLeapYear(2016) = %v, want %v", got, want)
		}
	}
	{
		got := IsLeapYear(2020)
		want := true
		if got != want {
			t.Errorf("IsLeapYear(2020) = %v, want %v", got, want)
		}

	}

	{
		got := IsLeapYear(2100)
		want := false
		if got != want {
			t.Errorf("IsLeapYear(2100) = %v, want %v", got, want)
		}
	}

	{
		got := IsLeapYear(2400)
		want := true
		if got != want {
			t.Errorf("IsLeapYear(2400) = %v, want %v", got, want)
		}
	}

	{
		got := IsLeapYear(2500)
		want := false
		if got != want {
			t.Errorf("IsLeapYear(2500) = %v, want %v", got, want)
		}
	}

	{
		got := IsLeapYear(3000)
		want := false
		if got != want {
			t.Errorf("IsLeapYear(3000) = %v, want %v", got, want)
		}
	}

	{
		got := IsLeapYear(4000)
		want := true
		if got != want {
			t.Errorf("IsLeapYear(4000) = %v, want %v", got, want)
		}
	}

}

/*
var _ = Describe("IsLeapYear", func() {
  Context("Fixed tests", func() {
    It("Leap years divisible by 4", func() {
      doTest(2020, true)
      doTest(1824, true)
      doTest(2152, true)
    })

    It("Leap years divisible by 400", func() {
      doTest(1600, true)
      doTest(2000, true)
      doTest(4000, true)
    })

    It("Non-leap years divisible by 100", func() {
      doTest(1800, false)
      doTest(1900, false)
      doTest(2100, false)
      doTest(2200, false)
    })

    It("Regular non-leap years", func() {
      doTest(1821, false)
      doTest(1942, false)
      doTest(2113, false)
      doTest(2254, false)
    })
  })
*/
