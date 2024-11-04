package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFin(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, FindShort("bitcoin take over the world maybe who knows perhaps"))
	assert.Equal(3, FindShort("turns out random test cases are easier than writing out basic ones"))
	assert.Equal(3, FindShort("lets talk about javascript the best language"))
	assert.Equal(1, FindShort("i want to travel the world writing code one day"))
	assert.Equal(2, FindShort("Lets all go on holiday somewhere very cold"))
}
