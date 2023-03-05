package main

import (
	"encoding/base64"
  
	"github.com/the-singularity-labs/Notary/pkg/signer"
	qrcode "github.com/skip2/go-qrcode"
  )

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {}

//export algoSign
func algoSign(contract string, seedPhrase string, msg string)  map[string]string {
	msg, err :=  signer.SignMessage(contract, seedPhrase, base64.StdEncoding.EncodeToString([]byte(msg)))
    if err != nil {
		return map[string]string{
			"error": err.Error(),
		}
    }

    png, err := qrcode.Encode(msg, qrcode.Medium, 256)
    if err != nil {
		return map[string]string{
			"error": err.Error(),
		}
    }
    base64QrCode:= "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
	return map[string]string{
		"signed_message": msg,
		"base64_qr_code": base64QrCode,
		"error": "",
	}
}