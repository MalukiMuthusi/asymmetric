package generatekeys_test

import (
	"math/big"
	"testing"

	"github.com/malukimuthusi/asymmetric/generatekeys"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrimeNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    big.Int
	}{
		{
			name: "happy case",
			n:    *big.NewInt(100),
		},
		{
			name: "generate large prime numbers",
			n:    *big.NewInt(1000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			n1, n2 := generatekeys.GeneratePrimeNumbers(tt.n)
			assert.NotEqualValues(t, n1, n2)
			assert.True(t, n1.Cmp(big.NewInt(1)) == 1)
			assert.True(t, n2.Cmp(big.NewInt(1)) == 1)
		})
	}

}

func TestProduct(t *testing.T) {
	tests := []struct {
		name string
		n1   big.Int
		n2   big.Int
		n    big.Int
	}{
		{
			name: "happy case",
			n1:   *big.NewInt(101_010),
			n2:   *big.NewInt(1_020_203),
			n:    *big.NewInt(103_050_705_030),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := generatekeys.Product(tt.n1, tt.n2)
			assert.Equal(t, tt.n, n)
		})
	}
}

func TestPhi(t *testing.T) {
	tests := []struct {
		name string
		p    big.Int
		q    big.Int
		phi  big.Int
	}{
		{
			name: "happy case",
			p:    *big.NewInt(991),
			q:    *big.NewInt(997),
			phi:  *big.NewInt(986040),
		},
		{
			name: "case 1",
			p:    *big.NewInt(2),
			q:    *big.NewInt(97),
			phi:  *big.NewInt(96),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phi := generatekeys.Phi(tt.p, tt.q)
			assert.Equal(t, tt.phi, phi)
		})
	}
}

func TestIsCoPrime(t *testing.T) {
	tests := []struct {
		name string
		a    big.Int
		b    big.Int
		is   bool
	}{
		{
			name: "happy case",
			a:    *big.NewInt(14),
			b:    *big.NewInt(5),
			is:   false,
		},
		{
			name: "case 2",
			a:    *big.NewInt(21),
			b:    *big.NewInt(14),
			is:   true,
		},

		{
			name: "case 3",
			a:    *big.NewInt(98),
			b:    *big.NewInt(5),
			is:   false,
		},

		{
			name: "case 4",
			a:    *big.NewInt(99),
			b:    *big.NewInt(3),
			is:   true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is := generatekeys.IsCoPrime(tt.a, tt.b)
			assert.Equal(t, tt.is, is)
		})

	}
}

func TestE(t *testing.T) {
	tests := []struct {
		name string
		n    big.Int
		p    big.Int
		q    big.Int
	}{
		{
			name: "happy case",
			p:    *big.NewInt(991),
			q:    *big.NewInt(997),
			n:    *big.NewInt(988027),
		},
		{
			name: "case 2",
			p:    *big.NewInt(991),
			q:    *big.NewInt(997),
			n:    *big.NewInt(988027),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := generatekeys.E(tt.n, tt.p, tt.q)

			half := new(big.Int).Div(&tt.n, big.NewInt(2))

			assert.Equal(t, -1, e.Cmp(half))
		})
	}
}

func TestD(t *testing.T) {
	tests := []struct {
		name string
		e    big.Int
		phi  big.Int
		d    big.Int
	}{
		{
			name: "case 1",
			e:    *big.NewInt(398309),
			phi:  *big.NewInt(986040),
			d:    *big.NewInt(241229),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := generatekeys.D(tt.e, tt.phi)

			de := new(big.Int).Mul(&d, &tt.e)

			mod := new(big.Int).Mod(de, &tt.phi)

			assert.EqualValues(t, mod.Int64(), int64(1))
		})
	}
}
