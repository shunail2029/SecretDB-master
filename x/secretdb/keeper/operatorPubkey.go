package keeper

import (
	"net/url"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

//
// Functions used by querier
//

func getOperatorPubkey(k Keeper) (res []byte, sdkError error) {
	res, err := codec.MarshalJSONIndent(k.cdc, url.PathEscape(string(types.OperatorPubkey.Bytes())))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
