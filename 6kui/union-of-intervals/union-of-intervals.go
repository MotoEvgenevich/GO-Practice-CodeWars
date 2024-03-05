package kata

import "sort"

func IntervalInsert(myl [][2]int, interval [2]int) [][2]int {
	result := make([][2]int, 0)
	for _, curr := range myl {
		if curr[1] < interval[0] || curr[0] > interval[1] {
			result = append(result, curr)
		} else {
			interval[0] = min(interval[0], curr[0])
			interval[1] = max(interval[1], curr[1])
		}
	}
	result = append(result, interval)
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Union of closed disjoint intervals

Write a function interval_insert which takes as input

a list myl of disjoint closed intervals with integer endpoints, sorted by increasing order of left endpoints
an interval interval
and returns the union of interval with the intervals in myl, as an array of disjoint closed intervals.

Terminology

A closed interval includes its endpoints. For example [0,5] means greater than or equal to 0 and less than or equal to 5. For more, refer to Wikipedia.

Two intervals are said to be disjoint if they have no element in common. Equivalently, disjoint intervals are intervals whose intersection is the empty interval. For example,
[1,3] and [4,6] are disjoint
[1,3] and [3,5] are not.
*/
