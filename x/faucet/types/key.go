package types

const (
	// ModuleName is the name of the module
	ModuleName = "faucet"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	RouterKey = ModuleName // this was defined in your key.go file

	QuerierRoute = ModuleName
)
