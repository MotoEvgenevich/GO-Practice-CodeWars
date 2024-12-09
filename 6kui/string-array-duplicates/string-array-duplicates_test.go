package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDup(t *testing.T) {
	{
		result := Dup([]string{"ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"})
		expect := []string{"codewars", "picaniny", "hubububo"}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Dup([]string{"abracadabra", "allottee", "assessee"})
		expect := []string{"abracadabra", "alote", "asese"}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Dup([]string{"kelless", "keenness"})
		expect := []string{"keles", "kenes"}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}

}

/*
   dotest([]string{"ccooddddddewwwaaaaarrrrsssss","piccaninny","hubbubbubboo"}, []string{"codewars","picaniny","hubububo"})
   dotest([]string{"abracadabra","allottee","assessee"}, []string{"abracadabra","alote","asese"})
   dotest([]string{"kelless","keenness"}, []string{"keles","kenes"})
*/
