package scrabble

import (
	"strings"
)

// Score receives a word and returns the scrabble score of the word
func Score(word string) int {
	var score int
	letters := strings.Split(word, "")
	for _, letter := range letters {
		scrab := strings.ToUpper(letter)
		if strings.Contains("A, E, I, O, U, L, N, R, S, T", scrab) {
			score++
		} else if strings.Contains("D, G", scrab) {
			score = score + 2
		} else if strings.Contains("B, C, M, P", scrab) {
			score = score + 3
		} else if strings.Contains("F, H, V, W, Y", scrab) {
			score = score + 4
		} else if strings.Contains("K", scrab) {
			score = score + 5
		} else if strings.Contains("J, X", scrab) {
			score = score + 8
		} else if strings.Contains("Q, Z", scrab) {
			score = score + 10
		}
	}
	return score
}
