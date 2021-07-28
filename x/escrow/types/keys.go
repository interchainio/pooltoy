package types

const (
	// ModuleName defines the module name
	ModuleName = "escrow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	//MemStoreKey = "mem_escrow"

	// this line is used by starport scaffolding # ibc/keys/name


	OfferPrefix ="offer-"
)

// this line is used by starport scaffolding # ibc/keys/port

var StartIndex = int64(0)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
