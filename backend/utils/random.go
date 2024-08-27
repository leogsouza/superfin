package utils

import "math/rand"

var alphabets string = "abcdefghijklmnopqrstuvwxyz"

func RandomString(r int) string {
	bits := []rune{}
	k := len(alphabets)

	for i := 0; i < r; i++ {
		idx := rand.Intn(k)
		bits = append(bits, rune(alphabets[idx]))
	}

	return string(bits)
}

func RandomEmail() string {
	return RandomString(8) + "@test.com"
}
