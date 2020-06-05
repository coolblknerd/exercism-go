package scrabble

import (
	"strings"
	"unicode"
)

// Score receives a word and returns the scrabble score of the word
func Score(word string) int {
	var score int
	letters := []rune(word)
	for _, letter := range letters {
		scrab := unicode.ToUpper(letter)
		switch {
		case strings.Contains("A, E, I, O, U, L, N, R, S, T", string(scrab)):
			score++
		case strings.Contains("D, G", string(scrab)):
			score = score + 2
		case strings.Contains("B, C, M, P", string(scrab)):
			score = score + 3
		case strings.Contains("F, H, V, W, Y", string(scrab)):
			score = score + 4
		case strings.Contains("K", string(scrab)):
			score = score + 5
		case strings.Contains("J, X", string(scrab)):
			score = score + 8
		case strings.Contains("Q, Z", string(scrab)):
			score = score + 10
		}
	}
	return score
}
