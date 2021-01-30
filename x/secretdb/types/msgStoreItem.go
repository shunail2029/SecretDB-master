package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var _ sdk.Msg = &MsgStoreItem{}

// MsgStoreItem is a message type to create item
type MsgStoreItem struct {
	Owner  sdk.AccAddress            `json:"owner" yaml:"owner"`
	Pubkey secp256k1.PubKeySecp256k1 `json:"pubkey" yaml:"pubkey"`
	Data   []byte                    `json:"data" yaml:"data"`
}

// NewMsgStoreItem returns new MsgStoreItem
func NewMsgStoreItem(owner sdk.AccAddress, pubkey secp256k1.PubKeySecp256k1, data []byte) MsgStoreItem {
	return MsgStoreItem{
		Owner:  owner,
		Pubkey: pubkey,
		Data:   data,
	}
}

// Route ...
func (msg MsgStoreItem) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgStoreItem) Type() string {
	return "CreateItem"
}

// GetSigners ...
func (msg MsgStoreItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgStoreItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgStoreItem) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
