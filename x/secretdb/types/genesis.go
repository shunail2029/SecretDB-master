package types

// GenesisState - all secretdb state that must be provided at genesis
type GenesisState struct {
	// FIXME: Fill out what is needed by the module for genesis
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* FIXME: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// FIXME: Fill out according to your genesis state
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// FIXME: Fill out according to your genesis state, these values will be initialized but empty
	}
}

// ValidateGenesis validates the secretdb genesis parameters
func ValidateGenesis(data GenesisState) error {
	// FIXME: Create a sanity check to make sure the state conforms to the modules needs
	return nil
}
