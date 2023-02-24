package main

import (
	"os"
	"fmt"
	"encoding/base64"

    "github.com/the-singularity-labs/Notary"
    
	"github.com/hoenirvili/skapt"
	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/flag"
)

func main() {
	app := skapt.Application{
		Name:        "Notary",
		Description: "Sign a message using an Algorand wallet for use in a TEAL dApp",
		Version:     "1.0.0",
		Handler: func(ctx *skapt.Context) error {
			var err error
			signerAcct := ctx.String("contract_account")
			seedPhrase := ctx.String("signer_seed_phrase")
			datab64 := ctx.String("datab64")

			
			msg, err := notary.SignMessage(signerAcct, seedPhrase, base64.StdEncoding.EncodeToString([]byte(datab64)))
			if err != nil {
				return err
			}

			fmt.Printf("Signed Message: %s\n", msg)

			return nil
		},
		Flags: flag.Flags{{
			Short: "a", Long: "contract_account",
			Description: "Countract address to sign for",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "s", Long: "signer_seed_phrase",
			Description: "Seed phrase of account",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "d", Long: "datab64",
			Description: "Data to be signed",
			Type:        argument.String,
			Required:	 true,
		}},
	}
	app.Exec(os.Args)
}
