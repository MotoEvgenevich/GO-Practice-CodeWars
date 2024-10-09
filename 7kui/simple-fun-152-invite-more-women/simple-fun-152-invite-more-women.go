package kata

func inviteMoreWomen(L []int) bool {
	menCount := 0
	womenCount := 0

	for _, person := range L {
		if person == 1 {
			menCount++
		} else if person == -1 {
			womenCount++
		}
	}

	return menCount > womenCount
}

/*
Task

King Arthur and his knights are having a New Years party. Last year Lancelot was jealous of Arthur,
because Arthur had a date and Lancelot did not, and they started a duel.

To prevent this from happening again, Arthur wants to make sure that there are at least as many women as men at this year's party.
He gave you a list of integers of all the party goers.

Arthur needs you to return true if he needs to invite more women or false if he is all set.

Input/Output

[input] integer array L ($a in PHP)
An array (guaranteed non-associative in PHP) representing the genders of the attendees, where -1 represents women and 1 represents men.

2 <= L.length <= 50

[output] a boolean value

true if Arthur need to invite more women, false otherwise.
*/
