package kata

func EachCons(arr []int, n int) [][]int {
	var result [][]int

	for i := 0; i <= len(arr)-n; i++ {
		sublist := make([]int, n)
		copy(sublist, arr[i:i+n])
		result = append(result, sublist)
	}

	return result
}

/*
DESCRIPTION:
Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n, like so:

each_cons([1,2,3,4], 2)
  #=> [[1,2], [2,3], [3,4]]

each_cons([1,2,3,4], 3)
  #=> [[1,2,3],[2,3,4]]

As you can see, the lists are cascading; ie, they overlap, but never out of order.
*/

/*
 TEST CASES:

 arr := []int{3, 5, 8, 13}
       Expect(EachCons(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
       Expect(EachCons(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
       Expect(EachCons(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
       Expect(EachCons([]int{}, 3)).To(BeEmpty())
*/
