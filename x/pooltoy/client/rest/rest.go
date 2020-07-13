package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers pooltoy-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/pooltoy/user", listUsersHandler(cliCtx, "pooltoy")).Methods("GET")
	r.HandleFunc("/pooltoy/user", createUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/faucet", faucetHandler(cliCtx)).Methods("POST")
}
