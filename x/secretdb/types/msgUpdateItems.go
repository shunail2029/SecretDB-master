package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateItems{}

// MsgUpdateItems is message type to set some items
type MsgUpdateItems struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter string         `json:"filter" yaml:"filter"`
	Update string         `json:"update" yaml:"update"`
}

// NewMsgUpdateItems returns new MsgUpdateItems
func NewMsgUpdateItems(owner sdk.AccAddress, filter string, update string) MsgUpdateItems {
	return MsgUpdateItems{
		Owner:  owner,
		Filter: filter,
		Update: update,
	}
}

// Route ...
func (msg MsgUpdateItems) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgUpdateItems) Type() string {
	return "SetItems"
}

// GetSigners ...
func (msg MsgUpdateItems) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgUpdateItems) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgUpdateItems) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
