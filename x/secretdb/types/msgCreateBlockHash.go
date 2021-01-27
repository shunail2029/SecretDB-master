package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBlockHash{}

// MsgCreateBlockHash ...
type MsgCreateBlockHash struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ChainID string         `json:"chainID" yaml:"chainID"`
	Height  int            `json:"height" yaml:"height"`
	Hash    []byte         `json:"hash" yaml:"hash"`
}

// NewMsgCreateBlockHash ...
func NewMsgCreateBlockHash(creator sdk.AccAddress, chainID string, height int, hash []byte) MsgCreateBlockHash {
	return MsgCreateBlockHash{
		Creator: creator,
		ChainID: chainID,
		Height:  height,
		Hash:    hash,
	}
}

// Route ...
func (msg MsgCreateBlockHash) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgCreateBlockHash) Type() string {
	return "CreateBlockHash"
}

// GetSigners ...
func (msg MsgCreateBlockHash) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgCreateBlockHash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreateBlockHash) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
