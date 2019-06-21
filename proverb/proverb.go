package proverb

import "fmt"

func forWant(want string, lost string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", want, lost)
}

func allFor(want string) string {
	return fmt.Sprintf("And all for the want of a %s.", want)
}

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverbs := []string{}

	for i, value := range rhyme {
		if value == rhyme[len(rhyme)-1] {
			proverbs = append(proverbs, allFor(rhyme[0]))
		} else {
			proverbs = append(proverbs, forWant(value, rhyme[i+1]))
		}
	}

	return proverbs
}
