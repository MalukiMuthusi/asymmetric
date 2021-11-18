package generatekeys

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kavehmz/prime"
)

// Generate two large prime numbers
func GeneratePrimeNumbers(n uint64) (uint64, uint64) {
	p := prime.Primes(n)
	l := len(p) - 1

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(l)
	n1 := p[i]

	r2 := rand.New(rand.NewSource(time.Now().UnixNano()))
	i2 := r2.Intn(l)
	n2 := p[i2]

	return n1, n2
}

func Product(n1, n2 uint64) uint64 {
	return n1 * n2
}

// Remove the co-primes and count the number of digits that remain
func Phi(p, q uint64) uint64 {

	return (p - 1) * (q - 1)

}

// choose e, a value less than the midpoint n
func E(n, p, q uint64) (uint64, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	s := r.Int63n(int64(n / 2))
	s2 := uint64(s)

	for i := s2; i > 0; i-- {
		if i%p != 0 && i%q != 0 {
			return uint64(i), nil
		}
	}

	return 0, fmt.Errorf("no value of E found")
}

func D(n, e, phi uint64) (uint64, error) {
	// d.e mod phi = 1

	for i := uint64(1); i < n; i++ {
		if (i*e)%phi == 1 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("%s", "failed to calculate D")

}
