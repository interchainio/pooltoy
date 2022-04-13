module github.com/interchainberlin/pooltoy

go 1.17

require (
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.6
)

require (
	github.com/kr/text v0.2.0 // indirect
	//github.com/cosmos/cosmos-sdk v0.45.1 //indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
