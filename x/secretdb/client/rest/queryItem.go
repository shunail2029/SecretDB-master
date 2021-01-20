package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func getItemHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filter := vars["filter"]
		pubkey := vars["pubkey"]
		sig := vars["sig"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get-item/%s/%s/%s", storeName, filter, pubkey, sig), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getItemsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filter := vars["filter"]
		pubkey := vars["pubkey"]
		sig := vars["sig"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get-items/%s/%s/%s", storeName, filter, pubkey, sig), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
