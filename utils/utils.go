package utils

import (
	"crypto/sha256"
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

func GenerateHash256(source string) string {
	h := sha256.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
