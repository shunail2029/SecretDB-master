package types

import (
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// flags
const (
	FlagOperatorName   = "operator-name"
	FlagKeyringBackend = "keyring-backend"
	FlagCLIHome        = "cli-home" // to use keyring
	FlagGas            = "gas"
	FlagChildCount     = "child-count"
	FlagChildURI       = "child-uri"
	FlagChildChainID   = "child-chainid"
)

// child chain params
var (
	OperatorName    string
	OperatorAddress sdk.AccAddress
	OperatorPubkey  crypto.PubKey
	KeyringBackend  string
	CLIHome         string
	Gas             uint64
	ChildCount      int
	ChildURIs       []string
	ChildChainIDs   []string
)

// SetParams ...
func SetParams(name, keyringBackend, cliHome string, gas uint64, count int, uris, chainIDs []string) error {
	var err error

	if name == "" {
		return errors.New("operator name must be specified")
	}
	OperatorName = name

	if keyringBackend == "" {
		return errors.New("keyring backend must be specified")
	}
	KeyringBackend = keyringBackend

	CLIHome = cliHome
	Gas = gas

	kb, err := keys.NewKeyring(sdk.KeyringServiceName(), KeyringBackend, CLIHome, os.Stdin)
	if err != nil {
		return err
	}
	info, err := kb.Get(OperatorName)
	if err != nil {
		return err
	}
	OperatorAddress = info.GetAddress()
	OperatorPubkey = info.GetPubKey()

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
