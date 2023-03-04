module github.com/the-singularity-labs/Notary/lib/wasm

go 1.20

replace github.com/the-singularity-labs/Notary/pkg/signer => ../../pkg/signer

require github.com/the-singularity-labs/Notary/pkg/signer v0.0.0-00010101000000-000000000000

require (
	github.com/algorand/go-algorand-sdk v1.24.0 // indirect
	github.com/algorand/go-codec/codec v1.1.8 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)
