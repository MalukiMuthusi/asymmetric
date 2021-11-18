package crypt

import (
	"math"
)

// Encrypt a message using the receivers public key and the senders private key
func Encrypt(msg int, e uint64, n uint64) uint64 {
	// E = msg ^ (privateKey) mod publicKey
	p := math.Pow(float64(msg), float64(e))
	pInt := uint64(p)

	c := pInt % n

	return c
}

// Decrypt a message using a receivers private key and the senders public key
func Decrypt(cypher uint64, d uint64, n uint64) uint64 {
	// M = cypher ^ privateKey mod publicKey

	p := math.Pow(float64(cypher), float64(d))
	pInt := uint64(p)

	t := pInt % n
	return t
}
