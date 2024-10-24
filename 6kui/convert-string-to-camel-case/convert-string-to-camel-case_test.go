package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	{
		result := ToCamelCase("the-stealth-warrior")
		expect := "theStealthWarrior"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ToCamelCase("the_Stealth_Warrior")
		expect := "theStealthWarrior"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
  	dotest("The_Stealth_Warrior","TheStealthWarrior")
    dotest("the-Stealth-Warrior","theStealthWarrior")
*/
