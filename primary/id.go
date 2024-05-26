package primary

import (
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("123456789011223344556677889900")

func IdRndm(length int)string{
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}