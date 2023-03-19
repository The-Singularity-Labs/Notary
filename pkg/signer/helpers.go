package signer

import (
	"crypto/ed25519"
)

// FromPrivateKey is a helper that converts an ed25519 private key to a
// human-readable mnemonic
func FromPrivateKey(sk ed25519.PrivateKey) (string, error) {
	seed := sk.Seed()
	return FromKey(seed)
}

// ToPrivateKey is a helper that converts a mnemonic directly to an ed25519
// private key
func ToPrivateKey(mnemonic string) (sk ed25519.PrivateKey, err error) {
	seedBytes, err := ToKey(mnemonic)
	if err != nil {
		return
	}
	return ed25519.NewKeyFromSeed(seedBytes), nil
}
