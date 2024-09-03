package kata

import (
	"fmt"
	"testing"
)

func TestBaumSweet(t *testing.T) {
	ch := make(chan int)
	go BaumSweet(ch)
	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}
}
