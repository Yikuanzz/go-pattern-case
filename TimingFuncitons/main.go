package main

import (
	"math/big"
	"time"

	"github.com/yikuanzz/go-pattern/TimingFuncitons/profile"
)

func BigIntFactorial(x *big.Int) *big.Int {
	defer profile.Duration(time.Now(), "BigIntFactorial")

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}

	return x.Set(y)
}

func main() {
	BigIntFactorial(big.NewInt(1000000000000000000))
}
