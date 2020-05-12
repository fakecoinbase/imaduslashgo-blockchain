package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/ripemd160"
)

const (
	checkSumLength = 4
	version        = byte(0x00)
)

//Wallet struct
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

//Address returns the address after hashing and converting to base58
func (wallet Wallet) Address() []byte {
	pubHash := PublicKeyHash(wallet.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checkSum := CheckSum(versionedHash)
	fullHash := append(versionedHash, checkSum...)

	address := Base58Encode(fullHash)
	return address

}

//NewKeyPair function returns the public and private key
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

//MakeWallet sets the public and private key for a wallet based on the user address
func MakeWallet() *Wallet {
	private, pub := NewKeyPair()

	wallet := Wallet{private, pub}

	return &wallet
}

//PublicKeyHash hashes the public key through a SHA-256
func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

//CheckSum hashes the public key twice through a SHA-256
func CheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checkSumLength]
}

//ValidateAddress validates the address passed to the wallet
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-checkSumLength:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checkSumLength]
	targetChecksum := CheckSum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}
