package kata

import (
	"sort"
	"strings"
)

type Guest struct {
	Name    string
	Surname string
}

type Guests []Guest

func Meeting(s string) string {
	// your code
	result := ""
	s = strings.ToUpper(s)

	var guests Guests

	for _, v := range strings.Split(s, ";") {
		var guest Guest
		guest.Name = strings.Split(v, ":")[0]
		guest.Surname = strings.Split(v, ":")[1]
		guests = append(guests, guest)
	}

	sort.Slice(guests, func(i, j int) bool {
		if guests[i].Surname == guests[j].Surname {
			return guests[i].Name < guests[j].Name
		}
		return guests[i].Surname < guests[j].Surname
	})

	for i := 0; i < len(guests); i++ {
		result += "(" + guests[i].Surname + ", " + guests[i].Name + ")"
	}
	return result
}

/*
John has invited some friends. His list is:

s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";
Could you make a program that

makes this string uppercase
gives it sorted in alphabetical order by last name.
When the last names are the same, sort them by first name. Last name and first name of a guest come in the result between parentheses separated by a comma.

So the result of function meeting(s) will be:

"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
It can happen that in two distinct families with the same family name two people have the same first name too.

Notes

You can see another examples in the "Sample tests".
*/
