package notary


import (
	"encoding/base64"
	"os"

	// "github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data/basics"
	// "github.com/algorand/go-algorand/data/transactions"
	// "github.com/algorand/go-algorand/data/transactions/logic"
	// "github.com/algorand/go-algorand/protocol"

	"crypto/ed25519"

)

const ProgramDataHashID string = "ProgData"

func SignMessage(signerAcct string, seedPhrase string, datab64 string) (signatureb64string, retErr error) {
	if signerAcct != "" {
		retErr = fmt.Errorf("Empty signerAcct is invalid")
		return
	}

	pK := NewKeyFromSeed([]byte(seedPhrase))



	// Otherwise, the contract address is the logic hash
	parsedAddr, err := basics.UnmarshalChecksumAddress(contractAddr)
	if err != nil {
		retErr = err
		return
	}

	/*
		* Next, fetch the data to sign
	*/

	dataToSign, err := base64.StdEncoding.DecodeString(datab64)
	if err != nil {
		retErr = err
		return
	}

	msg := append([]byte(ProgramDataHashID), append(parsedAddr[:], dataToSign...)...)

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