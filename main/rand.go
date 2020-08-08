package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	testPowRand()
}

func testPowRand() (err error) {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return err
	}
	randGen := rand.New(rand.NewSource(seed.Int64()))
	n := uint64(randGen.Uint64())
	fmt.Println(n)
	return nil
}
