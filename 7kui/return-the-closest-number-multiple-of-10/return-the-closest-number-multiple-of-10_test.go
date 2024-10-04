package kata

import "testing"

func TestClosestMultipleOf10(t *testing.T) {
	if got := ClosestMultipleOf10(54); got != 50 {
		t.Errorf("ClosestMultipleOf10(54) = %d; want 50", got)
	}
	if got := ClosestMultipleOf10(55); got != 60 {
		t.Errorf("ClosestMultipleOf10(55) = %d; want 60", got)
	}
	if got := ClosestMultipleOf10(22); got != 20 {
		t.Errorf("ClosestMultipleOf10(22) = %d; want 20", got)
	}
	if got := ClosestMultipleOf10(25); got != 30 {
		t.Errorf("ClosestMultipleOf10(25) = %d; want 30", got)
	}
	if got := ClosestMultipleOf10(37); got != 40 {
		t.Errorf("ClosestMultipleOf10(37) = %d; want 40", got)
	}
}
func TestClosestMultipleOf10v2(t *testing.T) {
	if got := ClosestMultipleOf10(54); got != 50 {
		t.Errorf("ClosestMultipleOf10(54) = %d; want 50", got)
	}
	if got := ClosestMultipleOf10(55); got != 60 {
		t.Errorf("ClosestMultipleOf10(55) = %d; want 60", got)
	}
	if got := ClosestMultipleOf10(22); got != 20 {
		t.Errorf("ClosestMultipleOf10(22) = %d; want 20", got)
	}
	if got := ClosestMultipleOf10(25); got != 30 {
		t.Errorf("ClosestMultipleOf10(25) = %d; want 30", got)
	}
	if got := ClosestMultipleOf10(37); got != 40 {
		t.Errorf("ClosestMultipleOf10(37) = %d; want 40", got)
	}
}

/*
   dotest(54, 50)
       dotest(55, 60)
	   Example input:

22
25
37
Expected output:

20
30
40
*/
