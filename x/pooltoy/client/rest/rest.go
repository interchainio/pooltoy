package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

//TODO: faucet request
// RegisterRoutes registers pooltoy-related REST handlers to a router
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/faucet", faucetHandler(cliCtx)).Methods("POST")
}
