package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := Zombie_shootout(3, 10, 10)
		expect := "You shot all 3 zombies."
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Zombie_shootout(100, 8, 200)
		expect := "You shot 16 zombies before being eaten: overwhelmed."
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}

	{
		result := Zombie_shootout(50, 10, 8)
		expect := "You shot 8 zombies before being eaten: ran out of ammo."
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Zombie_shootout(50, 10, 49)
		expect := "You shot 20 zombies before being eaten: overwhelmed."
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}

}

/*
It("Zombie_shootout(3, 10, 10)", func() { Expect(Zombie_shootout(3, 10, 10)).To(Equal("You shot all 3 zombies.")) })
  It("Zombie_shootout(100, 8, 200)", func() { Expect(Zombie_shootout(100, 8, 200)).To(Equal("You shot 16 zombies before being eaten: overwhelmed.")) })
  It("Zombie_shootout(50, 10, 8)", func() { Expect(Zombie_shootout(50, 10, 8)).To(Equal("You shot 8 zombies before being eaten: ran out of ammo.")) })
})
*/
