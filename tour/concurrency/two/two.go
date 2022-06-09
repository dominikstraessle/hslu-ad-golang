package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 10, 30, 2}

	c := make(chan int)
	go sum(s[:len(s)/3], c)
	go sum(s[len(s)/3:len(s)/3*2], c)
	go sum(s[len(s)/3*2:], c)
	x, y, z := <-c, <-c, <-c // receive from c

	fmt.Println(x, y, z, x+y+z)

	// check the result
	check := 0
	for _, i := range s {
		check += i
	}
	fmt.Println(check)
}
