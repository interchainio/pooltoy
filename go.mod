module github.com/interchainberlin/pooltoy

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.3
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.1
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kr/text v0.2.0 // indirect
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	github.com/tmdvs/Go-Emoji-Utils v1.1.0
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	google.golang.org/genproto v0.0.0-20210319143718-93e7006c17a6
	google.golang.org/grpc v1.36.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
