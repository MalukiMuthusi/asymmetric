package generatekeys

import (
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

func IsCoPrime(a, b uint64) bool {
	gcd := 1
	for i := uint64(1); i < b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd += 1
		}
	}
	return gcd == 1
}

// choose e, a value less than the midpoint n
func E(n, p, q uint64) (uint64, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	phi := Phi(p, q)

	for i := uint64(r.Int63n(int64(n / 2))); ; i = uint64(r.Int63n(int64(n / 2))) {
		if !IsCoPrime(i, p) || !IsCoPrime(i, q) || !IsCoPrime(i, phi) {
			return i, nil
		}
	}

}

func D(e, phi uint64) (uint64, error) {
	// d.e mod phi = 1

	d := uint64(1)
	for (d*e)%phi != 1 || d == e {
		d += 1
	}

	return d, nil

}
