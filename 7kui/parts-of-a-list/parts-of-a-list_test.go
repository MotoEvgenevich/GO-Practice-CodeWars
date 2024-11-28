package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartList(t *testing.T) {
	{
		result := PartList([]string{"I", "wish", "I", "hadn't", "come"})
		expect := "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)"
		assert.Equal(t, expect, result)
	}
	{
		result := PartList([]string{"cdIw", "tzIy", "xDu", "rThG"})
		expect := "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"
		assert.Equal(t, expect, result)
	}

}

/*
   dotest([]string{"I", "wish", "I", "hadn't", "come"},
       "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")
   dotest([]string{"cdIw", "tzIy", "xDu", "rThG"},
       "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)")
*/
