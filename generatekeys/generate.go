package generatekeys

import (
	"fmt"
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
func E(n, p, q int64) (int64, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	s := r.Int63n(n / 2)

	for i := s; i > 0; i-- {
		if i%p != 0 && i%q != 0 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no value of E found")
}

func D(e, phi int64) int64 {
	// d.e mod phi = 1

	d := 1 / (e % phi)

	return d
}
