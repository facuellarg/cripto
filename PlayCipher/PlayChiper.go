package PlayCipher

import (
	"bytes"
)

func encript(key string, msg string) string {
	var a bytes.Buffer
	moduloKey := len(key)
	moduloAlp := rune(26)

	for i, letter := range msg {
		letter -= 65
		letter = letter + rune(key[i%moduloKey]-65)
		a.WriteString(string((letter % moduloAlp) + 65))
	}
	return a.String()
}

func decript(key string, msg string) string {
	var a bytes.Buffer
	moduloKey := len(key)
	moduloAlp := rune(26)

	for i, letter := range msg {
		letter -= 65
		letter = letter - rune(key[i%moduloKey]-65)
		letter += 26
		a.WriteString(string((letter % moduloAlp) + 65))
	}
	return a.String()
}

//PalyCipherEncripter with 1 encript with another decript
func PlayCipherEncripter(key string, msg string, indicador int) string {
	if indicador == 1 {
		return encript(key, msg)
	}
	return decript(key, msg)
}
