package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteBlockHash{}

// MsgDeleteBlockHash ...
type MsgDeleteBlockHash struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ChainID string         `json:"chainID" yaml:"chainID"`
	Height  int            `json:"height" yaml:"height"`
}

// NewMsgDeleteBlockHash ...
func NewMsgDeleteBlockHash(creator sdk.AccAddress, chainID string, height int) MsgDeleteBlockHash {
	return MsgDeleteBlockHash{
		Creator: creator,
		ChainID: chainID,
		Height:  height,
	}
}

// Route ...
func (msg MsgDeleteBlockHash) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgDeleteBlockHash) Type() string {
	return "DeleteBlockHash"
}

// GetSigners ...
func (msg MsgDeleteBlockHash) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgDeleteBlockHash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgDeleteBlockHash) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
