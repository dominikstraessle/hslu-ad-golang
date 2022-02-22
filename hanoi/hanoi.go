package hanoi

import "fmt"

func hanoi(from, via, to string, n int) {
	if n == 1 {
		fmt.Println(fmt.Sprintf("move disk from %s to %s", from, to))
	} else {
		hanoi(from, to, via, n-1)
		fmt.Println(fmt.Sprintf("move disk from %s to %s", from, to))
		hanoi(via, from, to, n-1)
	}
}
