package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	shiftKey = shiftKey % 26
	shifted := make([]byte, len(plain))

	for i := 0; i < len(plain); i++ {
		if plain[i] >= 'a' && plain[i] <= 'z' {
			shifted[i] = (plain[i]-'a'+byte(shiftKey))%26 + 'a'
		} else if plain[i] >= 'A' && plain[i] <= 'Z' {
			shifted[i] = (plain[i]-'A'+byte(shiftKey))%26 + 'A'
		} else {
			shifted[i] = plain[i]
		}
	}

	return string(shifted)
}
