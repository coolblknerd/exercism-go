package strand

import (
	"strings"
)

// ToRNA transcribes DNA sequences to their RNA equivalent
func ToRNA(dna string) string {
	var transcription []string

	complement := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	for _, value := range strings.Split(dna, "") {
		transcription = append(transcription, complement[value])
	}

	return strings.Join(transcription, "")
}
