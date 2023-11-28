package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeyPair() (private string, public string, err error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	return string(crypto.FromECDSA(privKey)), string(crypto.FromECDSAPub(&privKey.PublicKey)), nil
}
