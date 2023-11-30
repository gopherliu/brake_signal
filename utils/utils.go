package utils

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeyPair() (private string, public string, err error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(crypto.FromECDSA(privKey)),
		hex.EncodeToString(crypto.FromECDSAPub(&privKey.PublicKey)),
		nil
}
