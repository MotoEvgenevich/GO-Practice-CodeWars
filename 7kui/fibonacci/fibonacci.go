package kata

func Fib(n int) int {
	fibSlice := []int{0, 1}
	fib := 0
	for i := 0; i <= n; i++ {
		fib = fibSlice[i] + fibSlice[i+1]
		fibSlice = append(fibSlice, fib)
	}
	return fibSlice[n]
}

/*
Create function fib that returns n'th element of Fibonacci sequence (classic programming task).
*/
