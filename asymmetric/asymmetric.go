package asymmetric

import (
	"math"
	"math/rand"

	"github.com/kavehmz/prime"
)

func GeneratePrimeNumbers() (uint64, uint64) {
	p := prime.Primes(10000)
	l := len(p)
	n := rand.Intn(l - 1)

	n1 := p[n-1]
	n2 := p[n1]

	return n1, n2
}

func Encrypt(msg int, privateKey uint64, publicKey uint64) int {
	// E = msg ^ (privateKey) mod publicKey
	n := math.Pow(float64(msg), float64(privateKey))
	c := uint64(n)

	e := c % publicKey

	return int(e)
}

func Decrypt(cypher int, privateKey uint64, publicKey uint64) int {
	// M = cypher ^ privateKey mod publicKey

	n := math.Pow(float64(cypher), float64(privateKey))
	c := uint64(n)

	e := c % publicKey
	return int(e)
}
