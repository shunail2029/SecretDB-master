package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// flags
const (
	FlagValidatorAccount = "validator-account"
	FlagKeyringBackend   = "keyring-backend"
	FlagCLIHome          = "cli-home" // to use keyring
	FlagChildCount       = "child-count"
	FlagChildURI         = "child-uri"
	FlagChildChainID     = "child-chainid"
)

// child chain params
var (
	ValidatorAccount sdk.AccAddress
	KeyringBackend   string
	CLIHome          string
	ChildCount       int
	ChildURIs        []string
	ChildChainIDs    []string
)

// SetChildParams ...
func SetChildParams(account, keyringBackend, cliHome string, count int, uris, chainIDs []string) error {
	if account == "" {
		return errors.New("validator account must be specified")
	}
	ValidatorAccount = sdk.AccAddress(account)

	if keyringBackend == "" {
		return errors.New("keyring backend must be specified")
	}
	KeyringBackend = keyringBackend

	CLIHome = cliHome

	if count == 0 {
		return errors.New("child count should be more than 0")
	}
	if count != len(uris) || count != len(chainIDs) {
		return errors.New("child count should be equal to length of child uris")
	}

	ChildCount = count
	ChildURIs = uris
	ChildChainIDs = chainIDs
	return nil
}
