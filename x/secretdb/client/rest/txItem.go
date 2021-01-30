package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type storeItemRequest struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Owner   sdk.AccAddress            `json:"owner"`
	Pubkey  secp256k1.PubKeySecp256k1 `json:"pubkey"`
	Data    []byte                    `json:"data"`
}

func storeItemHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req storeItemRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// check whether req.Data can be decoded as bson.M
		err := bson.UnmarshalExtJSON([]byte(req.Data), false, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgStoreItem(
			req.Owner,
			req.Pubkey,
			req.Data,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type updateItemRequest struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Owner   sdk.AccAddress            `json:"owner"`
	Pubkey  secp256k1.PubKeySecp256k1 `json:"pubkey"`
	Filter  []byte                    `json:"filter"`
	Update  []byte                    `json:"update"`
}

func updateItemHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateItemRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// check whether req.Filter and req.Update can be decoded as bson.M
		err := bson.UnmarshalExtJSON([]byte(req.Filter), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = bson.UnmarshalExtJSON([]byte(req.Update), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgUpdateItem(
			req.Owner,
			req.Pubkey,
			req.Filter,
			req.Update,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type updateItemsRequest struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Owner   sdk.AccAddress            `json:"owner"`
	Pubkey  secp256k1.PubKeySecp256k1 `json:"pubkey"`
	Filter  []byte                    `json:"filter"`
	Update  []byte                    `json:"update"`
}

func updateItemsHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateItemsRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// check whether req.Filter and req.Update can be decoded as bson.M
		err := bson.UnmarshalExtJSON([]byte(req.Filter), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = bson.UnmarshalExtJSON([]byte(req.Update), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgUpdateItems(
			req.Owner,
			req.Pubkey,
			req.Filter,
			req.Update,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteItemRequest struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Owner   sdk.AccAddress            `json:"owner"`
	Pubkey  secp256k1.PubKeySecp256k1 `json:"pubkey"`
	Filter  []byte                    `json:"filter"`
}

func deleteItemHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteItemRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// check whether req.Filter can be decoded as bson.M
		err := bson.UnmarshalExtJSON([]byte(req.Filter), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteItem(
			req.Owner,
			req.Pubkey,
			req.Filter,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteItemsRequest struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Owner   sdk.AccAddress            `json:"owner"`
	Pubkey  secp256k1.PubKeySecp256k1 `json:"pubkey"`
	Filter  []byte                    `json:"filter"`
}

func deleteItemsHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteItemsRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// check whether req.Filter can be decoded as bson.M
		err := bson.UnmarshalExtJSON([]byte(req.Filter), true, bson.M{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteItems(
			req.Owner,
			req.Pubkey,
			req.Filter,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
