package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// ########################### NOTES ###############################
// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// Good tutorial:
// Overview: https://www.youtube.com/watch?v=NmM9HA2MQGI
// Math: https://www.youtube.com/watch?v=Yjrfm_oRO0w
//	Big number: His video "n" = Our "p"
//  Alice's a: 1 <= a <= n; It is random
//	Bob's b: 1 <= b <= n; It is random

// Both calculate their Public Keys to share:
// A = g^a mod n <-- So this ends up between 0 and n
// B = g^b mod n

// This is called the "Discrete log problem" - given A or B, finding the right a or b is extremely hard since it could be any number between 1 and n thanks to modulus - we can not know what it is.

// Then, each of them take the Public Keys and raise it further
// Alice: B^a mod n -> (g^a modn)^b mod n
// Bob: A^b mod n -> (g^b modn)^a mod n
// So - both are ending up with g^ab mod n

// That's the shared secret key!

// If some hacker tried to calculate using the public component g^a and g^b, they'd end up with:
// g^a * g^b = g^(a+b)
// Not the same as our shared secret key!
// #################################################################

// Alice and Bob use Diffie-Hellman key exchange to share secrets.
// 1. They start with prime numbers
// 2. Pick private keys
// 3. Generate and share public keys
// 4. Generate a shared secret key

// 1. They start with prime numbers
// We return (a,A) = (private, public)
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}
// 2. Pick private keys
func PrivateKey(p *big.Int) *big.Int {
	return genBigRandNum(big.NewInt(2), p) // 2 <= a < p
}
// 3. Generate and share public keys
func PublicKey(private, p *big.Int, g int64) *big.Int {
	// https://pkg.go.dev/math/big#Int.Exp
	return new(big.Int).Exp(big.NewInt(g), private, p) // (g^private) % p
}
// 4. Generate a shared secret key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p) // (public)^private % p
}
// Utility functions
// Subtracts two "big" numbers
func SubtractBigBfromA(B, A *big.Int) *big.Int {
	return big.NewInt(0).Sub(B, A)
}
// Generates a random "big" number between min <= bg < max
func genBigRandNum(min, max *big.Int) *big.Int {
    bg := SubtractBigBfromA(max, min)
    n, err := rand.Int(rand.Reader, bg)
    if err != nil {
        panic(err)
    }
    return big.NewInt(0).Add(n, min)
}