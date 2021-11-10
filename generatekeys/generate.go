package generatekeys

import (
	"math/rand"

	"github.com/kavehmz/prime"
)

// Generate two large prime numbers
func GeneratePrimeNumbers() (uint64, uint64) {
	p := prime.Primes(10000)
	l := len(p)
	n := rand.Intn(l - 1)

	n1 := p[n-1]
	n2 := p[n1]

	return n1, n2
}

func Product(n1, n2 uint64) uint64 {
	return n1 * n2
}
