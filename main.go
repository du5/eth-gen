package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var T = 0

func main() {
	for {
		go func() {
			if T >= 2000 {
				return
			}
			T++
			pk, add := Gen()
			if add[:6] == "114514" || add[33:] == "1919810" {
				log.Printf("%s => %s", pk, add)
			}
			T--
		}()
	}
}

func Gen() (string, string) {
	pk, _ := crypto.GenerateKey()
	return hexutil.Encode(crypto.FromECDSA(pk))[2:], crypto.PubkeyToAddress(*pk.Public().(*ecdsa.PublicKey)).Hex()[2:]
}
