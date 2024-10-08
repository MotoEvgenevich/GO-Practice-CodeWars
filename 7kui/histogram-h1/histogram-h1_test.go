package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistogram(t *testing.T) {

	{
		result := Histogram([6]int{7, 3, 10, 1, 0, 5})
		except := "6|##### 5\n5|\n4|# 1\n3|########## 10\n2|### 3\n1|####### 7\n"
		assert.Equal(t, except, result, "result value doesn't equal expect value")

	}

	{
		result := Histogram([6]int{0, 0, 0, 0, 0, 0})
		except := "6|\n5|\n4|\n3|\n2|\n1|\n"
		assert.Equal(t, except, result, "result value doesn't equal expect value")
	}
}

/*
     It("Sample tests", func() {
        dotest([6]int{7,3,10,1,0,5}, "6|##### 5\n5|\n4|# 1\n3|########## 10\n2|### 3\n1|####### 7\n")
        dotest([6]int{0,0,0,0,0,0}, "6|\n5|\n4|\n3|\n2|\n1|\n")
     })
})
*/
