package kata

import "testing"

func TestSolution(t *testing.T) {

	{
		result := solution("abc", "c")
		expect := true
		if result != expect {
			t.Errorf("result doesn't equal expect")
		}
	}

	{
		result := solution("a", "z")
		expect := false
		if result != expect {
			t.Errorf("result doesn't equal expect")
		}
	}

	{
		result := solution("", "")
		expect := true
		if result != expect {
			t.Errorf("result doesn't equal expect")
		}
	}

	{
		result := solution("banana", "ana")
		expect := true
		if result != expect {
			t.Errorf("result doesn't equal expect")
		}
	}

}

/*

TEST IMPORT DATA:
Expect(solution("", "")).To(Equal(true))
Expect(solution(" ", "")).To(Equal(true))
Expect(solution("abc", "c")).To(Equal(true))
Expect(solution("banana", "ana")).To(Equal(true))
Expect(solution("a", "z")).To(Equal(false))
Expect(solution("", "t")).To(Equal(false))
Expect(solution("$a = $b + 1", "+1")).To(Equal(false))
Expect(solution("    ", "   ")).To(Equal(true))
Expect(solution("1oo", "100")).To(Equal(false))
*/
