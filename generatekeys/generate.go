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
	// phi = (p-1)*(q-1)

	n1 := new(big.Int).Sub(&p, big.NewInt(1))
	n2 := new(big.Int).Sub(&q, big.NewInt(1))

	phi := new(big.Int).Mul(n1, n2)

	return *phi
}

// check if a is co prime of b
func IsCoPrime(a, b big.Int) bool {
	one := big.NewInt(1)
	gcd := new(big.Int).GCD(nil, nil, &a, &b)
	return gcd.Cmp(one) == 0
}

// choose e, a value less than the midpoint n
func E(n, p, q big.Int) big.Int {
	// e < n/2
	one := big.NewInt(1)
	tw := big.NewInt(2)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	half := new(big.Int).Div(&n, tw)

	i := new(big.Int).Rand(r, half)

	phi := Phi(p, q)

	for {

		for j := i; j.Cmp(tw) == 1; j.Sub(j, one) {
			if IsCoPrime(*j, phi) {
				return *j
			}
		}

		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		i = new(big.Int).Rand(r, half)
	}

}

func D(e, phi big.Int) big.Int {
	// d.e mod phi = 1

	one := big.NewInt(1)

	d := big.NewInt(1)
	mod := new(big.Int)

	for mod.Cmp(one) != 0 {

		d = new(big.Int).Add(d, one)

		x := new(big.Int).Mul(d, &e)

		mod = new(big.Int).Mod(x, &phi)
	}

	return *d

}
