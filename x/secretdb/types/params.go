package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
	// FIXME: Define your default parameters
)

// Parameter store keys
var (
// FIXME: Define your keys for the parameter store
// KeyParamName          = []byte("ParamName")
)

// ParamKeyTable for secretdb module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for secretdb at genesis
type Params struct {
	// FIXME: Add your Paramaters to the Paramter struct
	// KeyParamName string `json:"key_param_name"`
}

// NewParams creates a new Params object
func NewParams( /* FIXME: Pass in the paramters*/ ) Params {
	return Params{
		// FIXME: Create your Params Type
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(`
	// FIXME: Return all the params as a string
	`)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		// FIXME: Pair your key with the param
		// params.NewParamSetPair(KeyParamName, &p.ParamName),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams( /* FIXME: Pass in your default Params */ )
}
