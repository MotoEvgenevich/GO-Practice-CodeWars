package uniq

func Uniq(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	result := []string{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			result = append(result, a[i])
		}
	}
	return result
}

/*
Implement a function which behaves like the uniq command in UNIX.

It takes as input a sequence and returns a sequence in which all duplicate elements following each other have been reduced to one instance.

Example:

["a", "a", "b", "b", "c", "a", "b", "c"]  =>  ["a", "b", "c", "a", "b", "c"]
*/
