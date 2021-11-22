package crypt

import (
	"math/big"
)

// Encrypt a message using the receivers public key and the senders private key
func Encrypt(msg, e, n big.Int) big.Int {
	// E = msg ^ (privateKey) mod publicKey
	var pp big.Int
	pp.Exp(&msg, &e, nil)

	var c big.Int
	c.Mod(&pp, &n)
	return c
}

// Decrypt a message using a receivers private key and the senders public key
func Decrypt(c, d, n big.Int) big.Int {
	// M = cypher ^ privateKey mod publicKey

	var pp big.Int
	pp.Exp(&c, &d, nil)

	var m big.Int
	m.Mod(&pp, &n)
	return m
}
