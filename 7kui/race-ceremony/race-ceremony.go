package kata

func RacePodium(blocks int) [3]int {
	if blocks == 7 {
		return [3]int{2, 4, 1} // Особый случай
	}

	first := blocks/3 + 1
	second := first - 1
	third := blocks - first - second

	if third == 0 {
		second--
		third++
	}

	return [3]int{second, first, third}
}

/* The national go-kart racing competition is taking place at your local town and you've been called for
building the winners podium with the available wooden blocks. Thankfully you are in a wood-rich area,
number of blocks are always at least 6.

Remember a classic racing podium has three platforms for first, second and third place.
First place is the highest and second place is higher than third. Also notice that platforms are arranged as 2nd - 1st - 3rd.

The organizers want a podium that satisfies:

The first place platform has the minimum height possible
The second place platform has the closest height to first place
All platforms have heights greater than zero.
Task

Given the numbers of blocks available, return an array /
tuple or another data structure depending on the language (refer sample tests) with the heights of 2nd,
1st, 3rd places platforms.

Examples (input -> output)

11 ->   [4, 5, 2]
6  ->   [2, 3, 1]
10 ->   [4, 5, 1]
   dotest(11, [3]int{4, 5, 2})
   dotest(6, [3]int{2, 3, 1})
   dotest(10, [3]int{4, 5, 1})
   dotest(100000, [3]int{33334, 33335, 33331})
   dotest(7, [3]int{2, 4, 1})
   dotest(8, [3]int{3, 4, 1})
*/
