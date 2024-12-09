package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeemojify(t *testing.T) {
	{
		result := Deemojify(":D :) :/  :D :) :|")
		expect := "hi"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Deemojify(";) >(  :D :) :D  :D :) ^.^  :D :) ^.^  :D :D :D  >:C >(  :D :D :(  :D :D :D  :D :D :/  :D :) ^.^  :D :) :)  >:C >:C")
		expect := "Hello world!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Deemojify(":)")
		expect := string("\x00")
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
Expect(Deemojify(":D :) :/  :D :) :|")).To(Equal("hi"))
    Expect(Deemojify(";) >(  :D :) :D  :D :) ^.^  :D :) ^.^  :D :D :D  >:C >(  :D :D :(  :D :D :D  :D :D :/  :D :) ^.^  :D :) :)  >:C >:C")).To(Equal("Hello world!"))
    Expect(Deemojify(":)")).To(Equal(string("\x00")))
*/
