package signer


import (
	"encoding/base32"
	"encoding/base64"
	"os"

	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions"
	"github.com/algorand/go-algorand/data/transactions/logic"
	"github.com/algorand/go-algorand/protocol"

)

func SignMessage(signerAcct string, seedPhrase string, datab64 string) (signatureb64string, retErr error) {
	if signerAcct != "" {
		retErr = fmt.Errorf("Empty signerAcct is invalid")
		return
	}

	// Create signature secrets from the seed
	var seed crypto.Seed
	copy(seed[:], seedPhrase)
	sec := crypto.GenerateSignatureSecrets(seed)


	var progHash crypto.Digest
	// Otherwise, the contract address is the logic hash
	parsedAddr, err := basics.UnmarshalChecksumAddress(contractAddr)
	if err != nil {
		retErr = err
		return
	}

	// Copy parsed address as program hash
	copy(progHash[:], parsedAddr[:])

	/*
		* Next, fetch the data to sign
		*/

	dataToSign, err := base64.StdEncoding.DecodeString(datab64)
	if err != nil {
		retErr = err
		return
	}

	/*
		* Sign the payload
		*/

	signature := sec.Sign(logic.Msg{
		ProgramHash: progHash,
		Data:        dataToSign,
	})

	// Always print signature to stdout
	signatureb64 = base64.StdEncoding.EncodeToString(signature[:])
	return signatureb64
}