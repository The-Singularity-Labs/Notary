package main

import (
  "syscall/js"
  "encoding/base64"

  "github.com/the-singularity-labs/Notary/pkg/signer"
  qrcode "github.com/skip2/go-qrcode"
)

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	js.Global().Set("algoSign", algoSignWrapper())
  <-make(chan bool)
}

func algoSignWrapper()  js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 3 {
			return wrap("", "Not enough arguments")
		}
		msg, err :=  signer.SignMessage(args[0].String(), args[1].String(), base64.StdEncoding.EncodeToString([]byte(args[2].String())))
    if err != nil {
      return wrap(msg, err.Error())
    }

    png, err := qrcode.Encode(msg, qrcode.Medium, 256)
    if err != nil {
      return wrap(map[string]interface{}{}, err.Error())
    }
    base64QrCode:= "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)

    return wrap(map[string]interface{}{"signed_message": msg, "base64_qr_code": base64QrCode}, "")
	})
}

func wrap(result interface{}, errString string) map[string]interface{} {
	return map[string]interface{}{
		"error":   errString,
		"data": result,
	}
}
