//Package faculty -> Week 01
package faculty

func factorialIter(n uint) uint {
	result := uint(1)
	for i := uint(0); i < n; i++ {
		result *= i
	}
	return result
}

func factorialRec(n uint) uint {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorialRec(n-1)
}
