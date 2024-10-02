package getvowelcount

func GetCount(str string) (count int) {

	for _, v := range str {
		if v == 97 || v == 101 || v == 105 || v == 111 || v == 117 {
			count += 1
		}
	}
	return count
}

/*  Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces. */

/*
should test that the solution returns the correct value			17
mcnwnot zonmly pjqikoig xwukdfj ft tthgie ue...					10
abracadabra														5
pmyhucyn lgbfkjn cpdln qodnhi hfwldqbq jdtnmtd kgw...			3
xcgfni q bzjcblj rxrfhql psufswl qju rphmv...					3
cyckvsrc cr ejdmugf yqbww iptjiwws sgxq							4
*/
