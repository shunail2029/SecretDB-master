package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateItem{}

// MsgUpdateItem is message type to set item
type MsgUpdateItem struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter string         `json:"filter" yaml:"filter"`
	Update string         `json:"update" yaml:"update"`
}

// NewMsgUpdateItem returns new MsgUpdateItem
func NewMsgUpdateItem(owner sdk.AccAddress, filter string, update string) MsgUpdateItem {
	return MsgUpdateItem{
		Owner:  owner,
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
