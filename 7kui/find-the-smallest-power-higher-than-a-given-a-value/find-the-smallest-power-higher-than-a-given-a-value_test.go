package kata

import (
	"testing"
)

func TestFindNextPower(t *testing.T) {
	result := FindNextPower(1245678, 5)
	if result != 1419857 {
		t.Errorf("Expected %v, got %v", 1419857, result)
	}
	result = FindNextPower(1245678, 6)
	if result != 1771561 {
		t.Errorf("Expected %v, got %v", 1771561, result)
	}
	{
		result := FindNextPower(12385, 3)
		if result != 13824 {
			t.Errorf("Expected %v, got %v", 13824, result)
		}
	}
}

/*
     It("Sample tests", func() {
        dotest(12385, 3, 13824)
        dotest(1245678, 5, 1419857)
        dotest(1245678, 6, 1771561)
     })
})
*/
