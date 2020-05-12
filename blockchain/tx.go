package blockchain

import (
	"bytes"

	"github.com/imadu/blockchain_app/wallet"
)

//TxInput struct
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

//TxOutput struct
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

//NewTXOutput creates a new transaction output and locks it
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

//UsesKey checkes if a public key exists in the wallet
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

//Lock locks the output against a given address
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

//IsLockedWithKey Checks is a users pubkey hash is the same as the transaction hash
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
