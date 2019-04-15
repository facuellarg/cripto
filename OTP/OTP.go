package otp

const MODULO = 26

//Encrypt the msg whit the key given
func Encrypt(msg string, key string) string {
	encrypted := make([]byte, len(msg))
	for i, c := range msg {
		c = c - 65
		newC := ((c + rune(key[i]-65)) % MODULO) + 65
		encrypted[i] = byte(newC)
	}
	return string(encrypted)
}

var keys map[string]bool

//KeyGenerator gen a new key for encrypt and drecipt
func KeyGenerator(msg string, result string) string {
	newKey := make([]byte, len(msg))
	for i, c := range msg {
		c = c - 65
		keyToEval := rune(result[i] - 65)
		var ck rune
		if keyToEval >= c {
			ck = keyToEval - c
		} else {
			ck = MODULO - c + keyToEval
		}
		newKey[i] = byte(ck + 65)
	}
	return string(newKey)
}

func cleanKeys() {
	keys = make(map[string]bool)
}
