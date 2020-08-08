package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_Te(t *testing.T) {
	s := "abc"
	s2 := hexutil.Bytes(s)
	fmt.Println(s2, string(s2))
}
