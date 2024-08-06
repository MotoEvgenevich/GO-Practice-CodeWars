package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzCuckooClock(t *testing.T) {
	{
		result := FizzBuzzCuckooClock("21:00")
		expected := "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
		assert.Equal(t, result, expected, "не равно в задании: 21:00")
	}
	{
		result := FizzBuzzCuckooClock("13:34")
		expected := "tick"
		assert.Equal(t, result, expected, "не равно в задании: 13:34")
	}
	{
		result := FizzBuzzCuckooClock("11:15")
		expected := "Fizz Buzz"
		assert.Equal(t, result, expected, "не равно в задании: 11:15")
	}
	{
		result := FizzBuzzCuckooClock("03:03")
		expected := "Fizz"
		assert.Equal(t, result, expected, "не равно в задании: 03:03")
	}
	{
		result := FizzBuzzCuckooClock("14:30")
		expected := "Cuckoo"
		assert.Equal(t, result, expected, "не равно в задании: 14:30")
	}
	{
		result := FizzBuzzCuckooClock("08:55")
		expected := "Buzz"
		assert.Equal(t, result, expected, "не равно в задании: 08:55")
	}
	{
		result := FizzBuzzCuckooClock("00:00")
		expected := "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
		assert.Equal(t, result, expected, "не равно в задании: 00:00")
	}
	{
		result := FizzBuzzCuckooClock("12:00")
		expected := "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
		assert.Equal(t, result, expected, "не равно в задании: 12:00")
	}

}

/*
Expect(FizzBuzzCuckooClock("13:34")).To(Equal("tick"))
     Expect(FizzBuzzCuckooClock("21:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
     Expect(FizzBuzzCuckooClock("11:15")).To(Equal("Fizz Buzz"))
     Expect(FizzBuzzCuckooClock("03:03")).To(Equal("Fizz"))
     Expect(FizzBuzzCuckooClock("14:30")).To(Equal("Cuckoo"))
     Expect(FizzBuzzCuckooClock("08:55")).To(Equal("Buzz"))
     Expect(FizzBuzzCuckooClock("00:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
     Expect(FizzBuzzCuckooClock("12:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
*/
