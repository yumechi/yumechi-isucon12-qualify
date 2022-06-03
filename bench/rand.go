package bench

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

var entropy *ulid.MonotonicEntropy

func init() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		// crypto/rand からReadできなかった場合の代替手段
		s = time.Now().UnixNano()
	}
	rand.Seed(s)

	entropy = ulid.Monotonic(rand.New(rand.NewSource(s)), 0)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
