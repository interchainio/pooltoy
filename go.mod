module github.com/interchainberlin/pooltoy

go 1.15

require (
	github.com/charleenfei/modules/incubator/faucet v0.0.0-20210317155022-f6e9bf473169 // indirect
	github.com/cosmos/cosmos-sdk v0.42.1
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b
	google.golang.org/grpc v1.36.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
