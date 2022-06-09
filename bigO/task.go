package bigO

import (
	"time"
)

func task(n int) {
	taskOne()
	taskOne()
	taskOne()
	taskOne()
	for i := 0; i < n; i++ {
		taskTwo()
		taskTwo()
		taskTwo()
		for j := 0; j < i; j++ {
			taskThree()
			taskThree()
		}
	}
}

func taskThree() {
	time.Sleep(time.Nanosecond * 100)
}

func taskTwo() {
	time.Sleep(time.Nanosecond * 100)
}

func taskOne() {
	time.Sleep(time.Nanosecond * 100)
}
