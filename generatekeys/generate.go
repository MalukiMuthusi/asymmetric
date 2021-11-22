package generatekeys

import (
	"math/big"
	"math/rand"
	"time"
)

// Generate two large prime numbers
func GeneratePrimeNumbers(n big.Int) (big.Int, big.Int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var p big.Int
	p.Rand(r, &n)

	for !p.ProbablyPrime(int(p.Int64())) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		p.Rand(r, &n)
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var q big.Int
	q.Rand(r, &n)
	for !q.ProbablyPrime(int(q.Int64())) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		q.Rand(r, &n)
	}

	return p, q
}

func Product(n1, n2 big.Int) big.Int {
	var p big.Int
	p.Mul(&n1, &n2)
	return p
}

// Remove the co-primes and count the number of digits that remain
func Phi(p, q big.Int) big.Int {
	var phi big.Int
	n1 := p.Sub(&p, big.NewInt(1))
	n2 := q.Sub(&p, big.NewInt(1))

	phi.Mul(n1, n2)

	return phi
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
func E(n, p, q big.Int) big.Int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var n2 big.Int
	n2.Rand(r, big.NewInt(0))

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
