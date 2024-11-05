package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPpg(t *testing.T) {
	{
		result := SumPpg(NBAPlayer{Team: "76ers", Ppg: 11.2}, NBAPlayer{Team: "bulls", Ppg: 20.2})
		expect := 31.4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumPpg(NBAPlayer{Team: "p1_team", Ppg: 20.2}, NBAPlayer{Team: "p2_team", Ppg: 2.6})
		expect := 22.8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumPpg(NBAPlayer{Team: "p3_team", Ppg: 2023.2}, NBAPlayer{Team: "p4_team", Ppg: 0})
		expect := 2023.2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumPpg(NBAPlayer{Team: "p5_team", Ppg: -5.8}, NBAPlayer{Team: "p6_team", Ppg: 0})
		expect := -5.8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumPpg(NBAPlayer{Team: "p7_team", Ppg: 0}, NBAPlayer{Team: "p8_team", Ppg: 0})
		expect := 0.0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
    iverson := NBAPlayer{ Team: "76ers", Ppg: 11.2 }
    jordan  := NBAPlayer{ Team: "bulls", Ppg: 20.2 }
   Expect(SumPpg(iverson,jordan)).To(BeNumerically("~",31.4,tol))
*/
/*
   player1 := NBAPlayer{ Team: "p1_team", Ppg: 20.2   }
   player2 := NBAPlayer{ Team: "p2_team", Ppg: 2.6    }
   player3 := NBAPlayer{ Team: "p3_team", Ppg: 2023.2 }
   player4 := NBAPlayer{ Team: "p4_team", Ppg: 0      }
   player5 := NBAPlayer{ Team: "p5_team", Ppg: -5.8   }
   Expect(SumPpg(player1, player2)).To(BeNumerically("~",22.8      ,tol))
   Expect(SumPpg(player3, player1)).To(BeNumerically("~",2043.4    ,tol))
   Expect(SumPpg(player3, player4)).To(BeNumerically("~",2023.2    ,tol))
   Expect(SumPpg(player4, player5)).To(BeNumerically("~",-5.8      ,tol))
   Expect(SumPpg(player5, player2)).To(BeNumerically("~",2.6 - 5.8 ,tol))
*/
