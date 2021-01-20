package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AccountAddressPrefix shows prefix of account's address
	AccountAddressPrefix = "cosmos"
)

var (
	// AccountPubKeyPrefix shows prefix of pubkey
	AccountPubKeyPrefix = AccountAddressPrefix + "pub"
	// ValidatorAddressPrefix shows prefix of validator's address
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	// ValidatorPubKeyPrefix shows prefix of validator's pubkey
	ValidatorPubKeyPrefix = AccountAddressPrefix + "valoperpub"
	// ConsNodeAddressPrefix ...
	ConsNodeAddressPrefix = AccountAddressPrefix + "valcons"
	// ConsNodePubKeyPrefix ...
	ConsNodePubKeyPrefix = AccountAddressPrefix + "valconspub"
)

// SetConfig ...
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
