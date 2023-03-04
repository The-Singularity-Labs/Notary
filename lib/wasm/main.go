package main

import (
  "github.com/the-singularity-labs/Notary/pkg/signer"
)

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {}


// This exports an algo_sign function.
// To make this function callable from JavaScript,
// we need to add the: "export algo_sign" comment above the function
//export algo_sign
func algo_sign(signerAcct string, seedPhrase string, datab64 string) string {
  msg, err := signer.SignMessage(signerAcct, seedPhrase, datab64)
  if err != nil {
    panic(err)
  }
  return msg;
}