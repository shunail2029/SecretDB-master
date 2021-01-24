package keeper

import (
	"errors"
	"net/url"
	"os"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptokeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/cosmos/cosmos-sdk/x/auth"
	authutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// if filter has "_owner", change it to owner, else add "_owner" to filter
func insertOwner(owner sdk.AccAddress, m bson.M) bson.M {
	if m == nil {
		m = make(bson.M)
	}
	m["_owner"] = owner.String()
	return m
}

func pathUnescape(path []string, k Keeper) ([]byte, secp256k1.PubKeySecp256k1, []byte, error) {
	msgStr, err := url.PathUnescape(path[0])
	if err != nil {
		return nil, secp256k1.PubKeySecp256k1{}, nil, err
	}
	pubkeyStr, err := url.PathUnescape(path[1])
	if err != nil {
		return nil, secp256k1.PubKeySecp256k1{}, nil, err
	}
	sigStr, err := url.PathUnescape(path[2])
	if err != nil {
		return nil, secp256k1.PubKeySecp256k1{}, nil, err
	}
	msg := []byte(msgStr)
	var pubkey secp256k1.PubKeySecp256k1
	k.cdc.UnmarshalBinaryBare([]byte(pubkeyStr), &pubkey) // XXX: only secp256k1 is accepted now
	sigBytes := []byte(sigStr)

	return msg, pubkey, sigBytes, nil
}

func calcRemainder(bz []byte, div int) int {
	res := 0
	for _, b := range bz {
		res = (256*res + int(b)) % div
	}
	return res
}

func sendQueryToChild(childNum int, query string) ([]byte, error) {
	if childNum > types.ChildCount {
		return nil, errors.New("childNum is incorrect")
	}

	chainID := types.ChildChainIDs[childNum]
	nodeURI := types.ChildURIs[childNum]

	// prepare CLIContext
	ctx := context.CLIContext{
		FromAddress: types.ValidatorAccount,
		ChainID:     chainID,
		NodeURI:     nodeURI,
	}

	// send query to child chain
	res, _, err := ctx.QueryWithData(query, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func sendTxToChild(childNum int, msgs []sdk.Msg) (sdk.TxResponse, error) {
	if childNum > types.ChildCount {
		return sdk.TxResponse{}, errors.New("childNum is incorrect")
	}

	chainID := types.ChildChainIDs[childNum]
	nodeURI := types.ChildURIs[childNum]

	// prepare CLIContext and TxBuilder
	ctx := context.CLIContext{
		FromAddress: types.ValidatorAccount,
		ChainID:     chainID,
		NodeURI:     nodeURI,
		FromName:    "validator",
	}
	kb, err := cryptokeys.NewKeyring(sdk.KeyringServiceName(), types.KeyringBackend, types.CLIHome, os.Stdin)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	txBldr, err := authutils.PrepareTxBuilder(
		auth.TxBuilder{}.WithChainID(chainID).WithKeybase(kb).WithTxEncoder(authutils.GetTxEncoder(codec.New())),
		ctx,
	)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(ctx.GetFromName(), keys.DefaultKeyPass, msgs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// broadcast tx to child chain
	res, err := ctx.BroadcastTxSync(txBytes)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return res, nil
}
