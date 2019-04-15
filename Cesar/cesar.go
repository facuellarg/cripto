package cesar

import (
	"bytes"
)

func encript(key rune, msg string, alp string) string {
	var a bytes.Buffer
	firtsValue := rune(alp[0])
	for _, i := range msg {
		i = i - firtsValue
		a.WriteString(string(alp[int(i+key)%len(alp)]))
	}
	return a.String()
}
func decript(key rune, msg string, alp string) string {
	var a bytes.Buffer
	firtsValue := rune(alp[0])
	for _, i := range msg {
		i = i - firtsValue - key + rune(len(alp))
		a.WriteString(string(alp[int(i)%len(alp)]))
	}
	return a.String()

}
