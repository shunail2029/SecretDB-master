package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// flags
const (
	FlagOperatorName    = "operator-name"
	FlagOperatorAddress = "operator-address"
	FlagKeyringBackend  = "keyring-backend"
	FlagCLIHome         = "cli-home" // to use keyring
	FlagGas             = "gas"
	FlagChildCount      = "child-count"
	FlagChildURI        = "child-uri"
	FlagChildChainID    = "child-chainid"
)

// child chain params
var (
	OperatorName    string
	OperatorAddress sdk.AccAddress
	KeyringBackend  string
	CLIHome         string
	Gas             uint64
	ChildCount      int
	ChildURIs       []string
	ChildChainIDs   []string
)

// SetParams ...
func SetParams(name, address, keyringBackend, cliHome string, gas uint64, count int, uris, chainIDs []string) error {
	var err error

	if name == "" {
		return errors.New("operator name must be specified")
	}
	OperatorName = name

	if address == "" {
		return errors.New("operator address must be specified")
	}
	OperatorAddress, err = sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	if keyringBackend == "" {
		return errors.New("keyring backend must be specified")
	}
	KeyringBackend = keyringBackend

	CLIHome = cliHome
	Gas = gas

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
