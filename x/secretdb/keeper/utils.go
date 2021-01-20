package keeper

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"
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
