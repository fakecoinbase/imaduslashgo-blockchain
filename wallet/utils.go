package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

//Base58Encode returns the pub key in base 58
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

//Base58Decode decodes a pub key in base 58
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
