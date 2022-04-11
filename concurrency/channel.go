package main

func channel() int {
	ch := make(chan int)
	go summarize(ch)
	return <-ch
}

func summarize(ch chan int) {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	ch <- sum
}
