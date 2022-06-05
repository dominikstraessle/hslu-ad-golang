//Package fibonacci -> Week 01
package fibonacci

func fiboRec(n int) int {
	if n <= 1 {
		return n
	}

	return fiboRec(n-2) + fiboRec(n-1)
}

func fiboIter(n int) int64 {
	if n <= 1 {
		return int64(n)
	}

	fibo2 := int64(0)
	fibo1 := int64(1)

	fibo := int64(0)

	for i := 2; i <= n; i++ {
		fibo = fibo2 + fibo1
		fibo2 = fibo1
		fibo1 = fibo
	}

	return fibo
}
