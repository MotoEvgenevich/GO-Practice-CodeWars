package kata

func CountZeros(n int) int {
	/* 	sum := big.NewInt(1)
	   	if n%2 == 0 {
	   		for i := 2; i <= n; {
	   			sum *= big.Int(uint64(i))
	   			i += 2
	   		}
	   		fmt.Println("odd sum", sum)
	   	}
	   	if n%2 != 0 {
	   		for i := 1; i <= n; {
	   			sum *= uint64(i)
	   			i += 2
	   		}
	   		fmt.Println("even sum", sum)
	   	} */

	return 0
}

/*
       Expect(CountZeros(8)).To(Equal(0))
       Expect(CountZeros(30)).To(Equal(3))
       Expect(CountZeros(487)).To(Equal(0))
       Expect(CountZeros(500)).To(Equal(62))

Define n!! as

n!! = 1 * 3 * 5 * ... * n if n is odd,

n!! = 2 * 4 * 6 * ... * n if n is even.

Hence 8!! = 2 * 4 * 6 * 8 = 384, there is no zero at the end. 30!! has 3 zeros at the end.

For a positive integer n, please count how many zeros are there at the end of n!!.

Example:

30!! = 2 * 4 * 6 * 8 * 10 * ... * 22 * 24 * 26 * 28 * 30
30!! = 42849873690624000 (3 zeros at the end) */
