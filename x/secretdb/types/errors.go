package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalid ...
	// code is 2 to avoid conflict with slave chain
	ErrInvalid = sdkerrors.Register(ModuleName, 2, "custom error message")
)
