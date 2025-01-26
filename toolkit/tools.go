package toolkit

import (
	"crypto/rand"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type that is used to instantiate the toolkit functions. Any variable of ths type will
// have access to all the functions in the toolkit package.

type Tools struct {
}

// RandomString generates a random string of length n. The string is generated using the randomStringSource
// as the source of characters. The function uses the crypto/rand package to generate the random numbers.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
