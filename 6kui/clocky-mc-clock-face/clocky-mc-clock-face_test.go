package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatTimeIsIt(t *testing.T) {
	{
		actual := WhatTimeIsIt(0)
		expect := "12:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(90)
		expect := "03:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(180)
		expect := "06:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(270)
		expect := "09:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(360)
		expect := "12:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(30)
		expect := "01:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(40)
		expect := "01:20"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(45)
		expect := "01:30"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(50)
		expect := "01:40"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := WhatTimeIsIt(60)
		expect := "02:00"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
}

/*
  {0, "12:00"},
  {90, "03:00"},
  {180, "06:00"},
  {270, "09:00"},
  {360, "12:00"},
  {30, "01:00"},
  {40, "01:20"},
  {45, "01:30"},
  {50, "01:40"},
  {60, "02:00"},
*/
