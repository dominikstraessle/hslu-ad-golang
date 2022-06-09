package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	p := New(2)

	for i := 0; i < 10; i++ {
		val := i
		p.Execute(func() {
			time.Sleep(time.Second)
			fmt.Println(val)
		})
		if i > 8 {
			p.Stop()
		}
	}
}
