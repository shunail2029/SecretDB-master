package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteItems{}

// MsgDeleteItems is a message type to delete some items
type MsgDeleteItems struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter string         `json:"filter" yaml:"filter"`
}

// NewMsgDeleteItems returns new MsgDeleteItems
func NewMsgDeleteItems(owner sdk.AccAddress, filter string) MsgDeleteItems {
	return MsgDeleteItems{
		Owner:  owner,
		Filter: filter,
	}
}

// Route ...
func (msg MsgDeleteItems) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgDeleteItems) Type() string {
	return "DeleteItems"
}

// GetSigners ...
func (msg MsgDeleteItems) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgDeleteItems) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgDeleteItems) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
