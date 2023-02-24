module github.com/the-singularity-labs/Notary/cmd/cli

go 1.20

replace github.com/the-singularity-labs/Notary => ../..

require (
	github.com/hoenirvili/skapt v0.0.0-20181026122304-fdaedd932adb
	github.com/the-singularity-labs/Notary v0.0.0-00010101000000-000000000000
)

require (
	github.com/algorand/go-algorand-sdk v1.24.0 // indirect
	github.com/algorand/go-codec/codec v1.1.8 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
