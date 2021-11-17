package crypt

import (
	"math"
)

// Encrypt a message using the receivers public key and the senders private key
func Encrypt(msg int, privateKey uint64, publicKey uint64) int {
	// E = msg ^ (privateKey) mod publicKey
	n := math.Pow(float64(msg), float64(privateKey))
	c := uint64(n)

	e := c % publicKey

	return int(e)
}

// Decrypt a message using a receivers private key and the senders public key
func Decrypt(cypher int, privateKey uint64, publicKey uint64) int {
	// M = cypher ^ privateKey mod publicKey

	n := math.Pow(float64(cypher), float64(privateKey))
	c := uint64(n)

	e := c % publicKey
	return int(e)
}
