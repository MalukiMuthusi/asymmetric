package generatekeys

import (
	"math/rand"
	"time"

	"github.com/kavehmz/prime"
)

// Generate two large prime numbers
func GeneratePrimeNumbers() (int64, int64) {
	p := prime.Primes(10000)
	l := len(p)
	n := rand.Intn(l - 1)

	n1 := p[n-1]
	n2 := p[n1]

	return int64(n1), int64(n2)
}

func Product(n1, n2 int64) int64 {
	return n1 * n2
}

// Remove the co-primes and count
func Phi(n, p, q int64) int64 {

	return (p - 1) * (q - 1)

}

// choose e, a value less than the midpoint n
func E(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return (r.Int63n(n / 2))
}

func D(e, phi int64) int64 {
	// d.e mod phi = 1

	d := 1 / (e % phi)
	
	return d
}
