package notary


import (
	"encoding/base32"
	"encoding/base64"
	"os"
	"crypto/ed25519"
	"crypto/sha512"
	// "github.com/algorand/go-algorand/crypto"
	// "github.com/algorand/go-algorand/data/basics"
	// "github.com/algorand/go-algorand/data/transactions"
	// "github.com/algorand/go-algorand/data/transactions/logic"
	// "github.com/algorand/go-algorand/protocol"
)

const ProgramDataHashID string = "ProgData"
const DigestSize = sha512.Size256

type CryptoDigest [DigestSize]byte

func SignMessage(signerAcct string, seedPhrase string, datab64 string) (signatureb64string, retErr error) {
	if signerAcct != "" {
		retErr = fmt.Errorf("Empty signerAcct is invalid")
		return
	}

	pK := NewKeyFromSeed([]byte(seedPhrase))

	var short CryptoDigest
	decoded, err := base32Encoder.DecodeString(address)
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
		retErr =err 
		return err
	}

	// Always print signature to stdout
	signatureb64 = base64.StdEncoding.EncodeToString(signature[:])
	return signatureb64
}