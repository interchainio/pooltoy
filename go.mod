module github.com/interchainberlin/pooltoy

go 1.13

require (
	// Oct 23rd, 2020 backports to go into 39.2
	github.com/cosmos/cosmos-sdk v0.39.2-0.20201023095550-c4f4d03f0f6e
	// github.com/cosmos/cosmos-sdk v0.38.4
	github.com/golang/mock v1.4.3 // indirect
	github.com/google/uuid v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/okwme/modules/incubator/faucet v0.0.0-20200719150004-606b92fc6e9c
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.8
	github.com/tendermint/tm-db v0.5.1
)

// replace github.com/okwme/modules/incubator/faucet => /Users/billy/GitHub.com/okwme/modules/incubator/faucet

// replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.6-0.20200802130156-46d1ad2d6210

// replace github.com/cosmos/cosmos-sdk v0.38.4 => /Users/billy/GitHub/cosmos/cosmos-sdk
