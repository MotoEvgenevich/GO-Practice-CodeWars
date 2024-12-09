package kata

func BabySharkLyrics() string {
	p := []string{"Baby", "Mommy", "Daddy", "Grandma", "Grandpa"}
	r, d := "", ", doo doo doo doo doo doo\n"
	for _, v := range p {
		r += v + " shark" + d + v + " shark" + d + v + " shark" + d + v + " shark!\n"
	}
	return r + "Let's go hunt" + d + "Let's go hunt" + d + "Let's go hunt" + d + "Let's go hunt!\nRun away,…\n"
}

/*
Baby Shark Lyrics

baby shark

Create a function, as short as possible, that returns this lyrics.
Your code should be less than 300 characters. Watch out for the three points at the end of the song.

Baby shark, doo doo doo doo doo doo
Baby shark, doo doo doo doo doo doo
Baby shark, doo doo doo doo doo doo
Baby shark!
Mommy shark, doo doo doo doo doo doo
Mommy shark, doo doo doo doo doo doo
Mommy shark, doo doo doo doo doo doo
Mommy shark!
Daddy shark, doo doo doo doo doo doo
Daddy shark, doo doo doo doo doo doo
Daddy shark, doo doo doo doo doo doo
Daddy shark!
Grandma shark, doo doo doo doo doo doo
Grandma shark, doo doo doo doo doo doo
Grandma shark, doo doo doo doo doo doo
Grandma shark!
Grandpa shark, doo doo doo doo doo doo
Grandpa shark, doo doo doo doo doo doo
Grandpa shark, doo doo doo doo doo doo
Grandpa shark!
Let's go hunt, doo doo doo doo doo doo
Let's go hunt, doo doo doo doo doo doo
Let's go hunt, doo doo doo doo doo doo
Let's go hunt!
Run away,…
*/
