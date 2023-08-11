package kata

import (
	"strconv"
)

func BonusTime(salary int, bonus bool) string {
	// Your code here
	result := ""
	if bonus {
		result = "£" + strconv.Itoa(salary*10)
		return result
	} else {
		result = "£" + strconv.Itoa(salary)
		return result
	}

}

/* It's bonus time in the big city! The fatcats are rubbing their paws in anticipation...
but who is going to make the most money?

Build a function that takes in two arguments (salary, bonus).
Salary will be an integer, and bonus a boolean.

If bonus is true, the salary should be multiplied by 10. If bonus is false,
the fatcat did not make enough money and must receive only his stated salary.

Return the total figure the individual will receive as a string prefixed with "£"
(= "\u00A3", JS, Go, Java, Scala, and Julia), "$" (C#, C++, Ruby,
Clojure, Elixir, PHP, Python, Haskell, and Lua) or "¥" (Rust). */

/* Expect(BonusTime(100, false)).To(Equal("£100"))
Expect(BonusTime(9, false)).To(Equal("£9"))
Expect(BonusTime(55000, false)).To(Equal("£55000"))
})
It("Adds a bonus if deserved", func() {
Expect(BonusTime(100, true)).To(Equal("£1000"))
Expect(BonusTime(14000, true)).To(Equal("£140000"))
}) */
