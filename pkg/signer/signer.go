package notary


import (
	"fmt"
	"encoding/base32"
	"encoding/base64"
	"crypto/ed25519"
	"crypto/sha512"

	"github.com/algorand/go-algorand-sdk/mnemonic"
)

const ProgramDataHashID string = "ProgData"
const DigestSize = sha512.Size256

var base32Encoder = base32.StdEncoding.WithPadding(base32.NoPadding)

type CryptoDigest [DigestSize]byte

func SignMessage(signerAcct string, seedPhrase string, datab64 string) (signatureb64 string, retErr error) {
	if signerAcct == "" {
		retErr = fmt.Errorf("Empty signerAcct is invalid")
		return
	}

	pK, err := mnemonic.ToPrivateKey(seedPhrase)
	if err != nil {
		retErr = err 
		return
	}

	var short CryptoDigest
	decoded, err := base32Encoder.DecodeString(signerAcct)
	if err != nil {
		retErr = err
		return
	}

	copy(short[:], decoded[:len(short)])

	// Otherwise, the contract address is the logic hash
	// parsedAddr, err := basics.UnmarshalChecksumAddress(contractAddr)
	// if err != nil {
	// 	retErr = err
	// 	return
	// }

	

	/*
		* Next, fetch the data to sign
	*/

	dataToSign, err := base64.StdEncoding.DecodeString(datab64)
	if err != nil {
		retErr = err
		return
	}

	msg := append([]byte(ProgramDataHashID), append(short[:], dataToSign...)...)

	/*
		* Sign the payload
	*/
	signature, err := pK.Sign(nil, msg, &ed25519.Options{Context: "notary_ed25519ctx"})
	if err != nil {
		retErr = err 
		return 
	}

	// Always print signature to stdout
	signatureb64 = base64.StdEncoding.EncodeToString(signature[:])
	return
}