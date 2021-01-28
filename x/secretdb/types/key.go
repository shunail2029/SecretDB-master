package types

const (
	// ModuleName is the name of the module
	ModuleName = "secretdb-master"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

// Prefixs
const (
	BlockHashPrefix = "BlockHash-"
)
