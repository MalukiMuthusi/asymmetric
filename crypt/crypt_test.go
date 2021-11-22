package crypt_test

import (
	"math/big"
	"testing"

	"github.com/malukimuthusi/asymmetric/crypt"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		msg  big.Int
		e    big.Int
		n    big.Int
		c    big.Int
	}{
		{
			name: "happy case",
			msg:  *big.NewInt(18),
			e:    *big.NewInt(28477),
			n:    *big.NewInt(988027),
			c:    *big.NewInt(387499),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := crypt.Encrypt(tt.msg, tt.e, tt.n)
			assert.EqualValues(t, tt.c, c)
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name string
		c    big.Int
		d    big.Int
		n    big.Int
		m    big.Int
	}{
		{
			name: "happy case",
			c:    *big.NewInt(387499),
			n:    *big.NewInt(988027),
			d:    *big.NewInt(769213),
			m:    *big.NewInt(18),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := crypt.Decrypt(tt.c, tt.d, tt.n)
			assert.EqualValues(t, tt.m, m)
		})
	}
}
