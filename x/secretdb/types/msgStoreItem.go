package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStoreItem{}

// MsgStoreItem is a message type to create item
type MsgStoreItem struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Data  string         `json:"data" yaml:"data"`
}

// NewMsgStoreItem returns new MsgStoreItem
func NewMsgStoreItem(owner sdk.AccAddress, data string) MsgStoreItem {
	return MsgStoreItem{
		Owner: owner,
		Data:  data,
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
