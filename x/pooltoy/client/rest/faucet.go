package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type claimReq struct {
	Address string
}

func faucetHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var claim claimReq
		decoder := json.NewDecoder(r.Body)
		decoderErr := decoder.Decode(&claim)
		if decoderErr != nil {
			panic(decoderErr)
		}
		// make sure address is bech32
		_, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, claim.Address)
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", err))
		}
		cmd := exec.Command("pooltoycli", "tx", "send", "me", claim.Address, "1token", "-y")
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s", err))
		}
		rest.PostProcessResponse(w, cliCtx, claim.Address)
	}
}
