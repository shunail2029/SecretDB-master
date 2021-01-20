package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteItem{}

// MsgDeleteItem is a message type to delete item
type MsgDeleteItem struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter string         `json:"filter" yaml:"filter"`
}

// NewMsgDeleteItem returns new MsgDeleteItem
func NewMsgDeleteItem(owner sdk.AccAddress, filter string) MsgDeleteItem {
	return MsgDeleteItem{
		Owner:  owner,
		Filter: filter,
	}
}

// Route ...
func (msg MsgDeleteItem) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgDeleteItem) Type() string {
	return "DeleteItem"
}

// GetSigners ...
func (msg MsgDeleteItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes ...
func (msg MsgDeleteItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgDeleteItem) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
