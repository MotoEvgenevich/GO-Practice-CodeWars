/* Complete the function which converts a binary number (given as a string) to a decimal number.
TESTS:
       Expect(BinToDec("0")).To(Equal(0))
       Expect(BinToDec("1")).To(Equal(1))
       Expect(BinToDec("10")).To(Equal(2))
       Expect(BinToDec("11")).To(Equal(3))
       Expect(BinToDec("101010")).To(Equal(42))
       Expect(BinToDec("1001001")).To(Equal(73))
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(BinToDec("1001001"))
	fmt.Println(BinToDec("1101111111000000101110")) //3665966

}

func BinToDec(bin string) int {
	decimalNum := 0
	var err error
	var digit int

	for i := 0; i < len(bin); i++ {
		if digit, err = strconv.Atoi(string(bin[len(bin)-1-i])); err != nil {
			fmt.Println("Error in binary digit conversion: ", err)
			return -1
		}
		fmt.Println("decimal num before:", decimalNum)
		decimalNum = decimalNum + digit*int(math.Pow(2, float64(i)))
		fmt.Println(decimalNum, "+", digit, "* 2", "**", i)
	}

	return decimalNum
}

func BinToDecV2(bin string) int {
	r, _ := strconv.ParseInt(bin, 2, 64)
	return int(r)
}

func BinToDecV3(bin string) int {
	n := 0
	for _, r := range bin {
		n *= 2
		n += int(r - '0')
	}
	return n
}
