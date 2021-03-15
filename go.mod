module github.com/interchainberlin/pooltoy

go 1.13

require (
	// Oct 23rd, 2020 backports to go into 39.2
	github.com/cosmos/cosmos-sdk v0.41.3
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/okwme/modules/incubator/faucet v0.0.0-20200719150004-606b92fc6e9c
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.7
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b
	google.golang.org/grpc v1.36.0
	rsc.io/quote/v3 v3.1.0 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

// replace github.com/okwme/modules/incubator/faucet => /Users/billy/GitHub.com/okwme/modules/incubator/faucet

// replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.6-0.20200802130156-46d1ad2d6210

// replace github.com/cosmos/cosmos-sdk v0.38.4 => /Users/billy/GitHub/cosmos/cosmos-sdk
