package main

import (
  "syscall/js"
  "encoding/base64"

  "github.com/the-singularity-labs/Notary/pkg/signer"
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
    return wrap(msg, "")
	})
}

func wrap(result string, errString string) map[string]interface{} {
	return map[string]interface{}{
		"error":   errString,
		"data": result,
	}
}
