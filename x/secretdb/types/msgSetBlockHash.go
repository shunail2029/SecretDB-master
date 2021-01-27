package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetBlockHash{}

// MsgSetBlockHash ...
type MsgSetBlockHash struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ChainID string         `json:"chainId" yaml:"chainId"`
	Height  int64          `json:"height" yaml:"height"`
	Hash    []byte         `json:"hash" yaml:"hash"`
}

// NewMsgSetBlockHash ...
func NewMsgSetBlockHash(creator sdk.AccAddress, chainID string, height int64, hash []byte) MsgSetBlockHash {
	return MsgSetBlockHash{
		Creator: creator,
		ChainID: chainID,
		Height:  height,
		Hash:    hash,
	}
}

// Route ...
func (msg MsgSetBlockHash) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgSetBlockHash) Type() string {
	return "SetBlockHash"
}

// GetSigners ...
func (msg MsgSetBlockHash) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgSetBlockHash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgSetBlockHash) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
