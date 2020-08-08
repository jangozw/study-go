package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	file := "../go.sum"
	inFile, _ := os.Open(file)
	sha1h := sha1.New()
	io.Copy(sha1h, inFile)
	fmt.Printf("%x   %s\n", sha1h.Sum([]byte("")), file)

	sha256h := sha256.New()
	io.Copy(sha256h, inFile)
	r := sha256h.Sum([]byte(""))
	fmt.Println(r)
	fmt.Printf("%x\n", r)
	str2 := string(r)

	fmt.Println(str2, len(str2))

	var hashInt *big.Int
	hashInt.SetBytes(r)
	var y *big.Int
	hashInt.Cmp(y)

	// rp160 := crypto.RIPEMD160.New()
	// str3 := rp160.Sum([]byte(str2))
	// fmt.Println(str3)
}
