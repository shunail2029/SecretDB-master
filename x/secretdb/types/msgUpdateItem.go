package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var _ sdk.Msg = &MsgUpdateItem{}

// MsgUpdateItem is message type to set item
type MsgUpdateItem struct {
	Owner  sdk.AccAddress            `json:"owner" yaml:"owner"`
	Pubkey secp256k1.PubKeySecp256k1 `json:"pubkey" yaml:"pubkey"`
	Filter []byte                    `json:"filter" yaml:"filter"`
	Update []byte                    `json:"update" yaml:"update"`
}

// NewMsgUpdateItem returns new MsgUpdateItem
func NewMsgUpdateItem(owner sdk.AccAddress, pubkey secp256k1.PubKeySecp256k1, filter []byte, update []byte) MsgUpdateItem {
	return MsgUpdateItem{
		Owner:  owner,
		Pubkey: pubkey,
		Filter: filter,
		Update: update,
	}
}

// Route ...
func (msg MsgUpdateItem) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgUpdateItem) Type() string {
	return "SetItem"
}

// GetSigners ...
func (msg MsgUpdateItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgUpdateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgUpdateItem) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
