package main

import (
	"strconv"
)

func SubtractSum(n int) string {
	var fruits = []string{
		"kiwi", "pear", "kiwi", "banana", "melon", "banana", "melon",
		"pineapple", "apple", "pineapple", "cucumber", "pineapple",
		"cucumber", "orange", "grape", "orange", "grape", "apple",
		"grape", "cherry", "pear", "cherry", "pear", "kiwi",
		"banana", "kiwi", "apple", "melon", "banana", "melon",
		"pineapple", "melon", "pineapple", "cucumber", "orange",
		"apple", "orange", "grape", "orange", "grape", "cherry",
		"pear", "cherry", "pear", "apple", "pear", "kiwi", "banana",
		"kiwi", "banana", "melon", "pineapple", "melon", "apple",
		"cucumber", "pineapple", "cucumber", "orange", "cucumber",
		"orange", "grape", "cherry", "apple", "cherry", "pear",
		"cherry", "pear", "kiwi", "pear", "kiwi", "banana", "apple",
		"banana", "melon", "pineapple", "melon", "pineapple",
		"cucumber", "pineapple", "cucumber", "apple", "grape",
		"orange", "grape", "cherry", "grape", "cherry", "pear",
		"cherry", "apple", "kiwi", "banana", "kiwi", "banana",
		"melon", "banana", "melon", "pineapple", "apple", "pineapple",
	}

	for n >= 10 && n < 10000 {
		sum := 0
		for _, digit := range strconv.Itoa(n) {
			val, _ := strconv.Atoi(string(digit))
			sum += val
		}
		n -= sum
		if n <= len(fruits) {
			return fruits[n-1]
		}
	}

	return "Not in the list"
}

/* type fruitsAndVegetables struct {
	number int
	name   string
}

type FruitsList struct {
	Fruits [100]fruitsAndVegetables
}

var fruits = FruitsList{
	Fruits: [100]fruitsAndVegetables{
		{number: 1, name: "kiwi"},
		{number: 2, name: "pear"},
		{number: 3, name: "kiwi"},
		{number: 5, name: "melon"},
		{number: 6, name: "banana"},
		{number: 7, name: "melon"},
		{number: 8, name: "pineapple"},
		{number: 9, name: "apple"},
		{number: 10, name: "pineapple"},
		{number: 11, name: "cucumber"},
		{number: 12, name: "pineapple"},
		{number: 13, name: "cucumber"},
		{number: 14, name: "orange"},
		{number: 15, name: "grape"},
		{number: 16, name: "orange"},
		{number: 17, name: "grape"},
		{number: 18, name: "apple"},
		{number: 19, name: "grape"},
		{number: 20, name: "cherry"},
		{number: 21, name: "pear"},
		{number: 22, name: "cherry"},
		{number: 23, name: "pear"},
		{number: 24, name: "kiwi"},
		{number: 25, name: "banana"},
		{number: 26, name: "kiwi"},
		{number: 27, name: "apple"},
		{number: 28, name: "melon"},
		{number: 29, name: "banana"},
		{number: 30, name: "melon"},
		{number: 31, name: "pineapple"},
		{number: 32, name: "melon"},
		{number: 33, name: "pineapple"},
		{number: 34, name: "cucumber"},
		{number: 35, name: "orange"},
		{number: 36, name: "apple"},
		{number: 37, name: "orange"},
		{number: 38, name: "grape"},
		{number: 39, name: "orange"},
		{number: 40, name: "grape"},
		{number: 41, name: "cherry"},
		{number: 42, name: "pear"},
		{number: 43, name: "cherry"},
		{number: 44, name: "pear"},
		{number: 45, name: "apple"},
		{number: 46, name: "pear"},
		{number: 47, name: "kiwi"},
		{number: 48, name: "banana"},
		{number: 49, name: "kiwi"},
		{number: 50, name: "banana"},
		{number: 51, name: "melon"},
		{number: 52, name: "pineapple"},
		{number: 53, name: "melon"},
		{number: 54, name: "apple"},
		{number: 55, name: "cucumber"},
		{number: 56, name: "pineapple"},
		{number: 57, name: "cucumber"},
		{number: 58, name: "orange"},
		{number: 59, name: "cucumber"},
		{number: 60, name: "orange"},
		{number: 61, name: "grape"},
		{number: 62, name: "cherry"},
		{number: 63, name: "apple"},
		{number: 64, name: "cherry"},
		{number: 65, name: "pear"},
		{number: 66, name: "cherry"},
		{number: 67, name: "pear"},
		{number: 68, name: "kiwi"},
		{number: 69, name: "pear"},
		{number: 70, name: "kiwi"},
		{number: 71, name: "banana"},
		{number: 72, name: "apple"},
		{number: 73, name: "banana"},
		{number: 74, name: "melon"},
		{number: 75, name: "pineapple"},
		{number: 76, name: "melon"},
		{number: 77, name: "pineapple"},
		{number: 78, name: "cucumber"},
		{number: 79, name: "pineapple"},
		{number: 80, name: "cucumber"},
		{number: 81, name: "apple"},
		{number: 82, name: "grape"},
		{number: 83, name: "orange"},
		{number: 84, name: "grape"},
		{number: 85, name: "cherry"},
		{number: 86, name: "grape"},
		{number: 87, name: "cherry"},
		{number: 88, name: "pear"},
		{number: 89, name: "cherry"},
		{number: 90, name: "apple"},
		{number: 91, name: "kiwi"},
		{number: 92, name: "banana"},
		{number: 93, name: "kiwi"},
		{number: 94, name: "banana"},
		{number: 95, name: "melon"},
		{number: 96, name: "banana"},
		{number: 97, name: "melon"},
		{number: 98, name: "pineapple"},
		{number: 99, name: "apple"},
		{number: 100, name: "pineapple"},
	},
}

func main() {
	fmt.Println("------------")
	//fmt.Println(fruits.Fruits[1])
	fmt.Println(SubtractSum(10))

} */

/*
	 func SubtractSum(n int) string {
		result := n
		digits := numberToSlice(n)
		sum := 0
		for result > 100 {
			for _, v := range digits {
				sum += v
			}
			result -= sum
			sum = 0
		}
		fmt.Println(result)
		return fruits.Fruits[result-1].name
	}
*/
/* func SubtractSum(n int) string {
	result := n
	for result > 100 {
		digits := numberToSlice(result)
		sum := 0
		for _, v := range digits {
			sum += v
		}
		result -= sum
	}

	if result <= 0 || result > len(fruits.Fruits) {
		return "Index out of range"
	}
	//fmt.Println(result)
	fmt.Println(fmt.Println(fruits.Fruits[3].name))
	return fruits.Fruits[result-1].name // Индексация с 0
}

func numberToSlice(number int) []int {
	// Convert the number to a string
	str := strconv.Itoa(number)

	// Create a slice to store the digits
	var digits []int

	// Iterate through each character in the string
	for _, char := range str {
		// Convert the character to a number
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			return nil
		}

		// Append the digit to the slice
		digits = append(digits, digit)
	}
	// Return the slice of digits
	return digits
} */

/* Subtract the sum

NOTE! This kata can be more difficult than regular 8-kyu katas (lets say 7 or 6 kyu)

Complete the function which get an input number n such that n >= 10 and n < 10000, then:

Sum all the digits of n.
Subtract the sum from n, and it is your new n.
If the new n is in the list below return the associated fruit, otherwise return back to task 1.
Example

n = 325
sum = 3+2+5 = 10
n = 325-10 = 315 (not in the list)
sum = 3+1+5 = 9
n = 315-9 = 306 (not in the list)
sum = 3+0+6 = 9
n =306-9 = 297 (not in the list)
.
.
.
...until you find the first n in the list below.

There is no preloaded code to help you. This is not about coding skills; think before you code

1-kiwi
2-pear
3-kiwi
4-banana
5-melon
6-banana
7-melon
8-pineapple
9-apple
10-pineapple
11-cucumber
12-pineapple
13-cucumber
14-orange
15-grape
16-orange
17-grape
18-apple
19-grape
20-cherry
21-pear
22-cherry
23-pear
24-kiwi
25-banana
26-kiwi
27-apple
28-melon
29-banana
30-melon
31-pineapple
32-melon
33-pineapple
34-cucumber
35-orange
36-apple
37-orange
38-grape
39-orange
40-grape
41-cherry
42-pear
43-cherry
44-pear
45-apple
46-pear
47-kiwi
48-banana
49-kiwi
50-banana
51-melon
52-pineapple
53-melon
54-apple
55-cucumber
56-pineapple
57-cucumber
58-orange
59-cucumber
60-orange
61-grape
62-cherry
63-apple
64-cherry
65-pear
66-cherry
67-pear
68-kiwi
69-pear
70-kiwi
71-banana
72-apple
73-banana
74-melon
75-pineapple
76-melon
77-pineapple
78-cucumber
79-pineapple
80-cucumber
81-apple
82-grape
83-orange
84-grape
85-cherry
86-grape
87-cherry
88-pear
89-cherry
90-apple
91-kiwi
92-banana
93-kiwi
94-banana
95-melon
96-banana
97-melon
98-pineapple
99-apple
100-pineapple */
