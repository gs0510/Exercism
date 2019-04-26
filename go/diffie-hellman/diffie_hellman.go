package diffiehellman

import "math/big"
import "math/rand"

var a big.Int
var b big.Int

func PrivateKey(p *big.Int) *big.Int {
	key := big.NewInt()
	randomGenerator := rand.New(rand.NewSource(44))
	key.Rand(randomGenerator, p)
	return key
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return p
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return p
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	return p, p
}
