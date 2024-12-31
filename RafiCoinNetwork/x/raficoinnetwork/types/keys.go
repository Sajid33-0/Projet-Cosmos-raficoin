package types

const (
	// ModuleName defines the module name
	ModuleName = "raficoinnetwork"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_raficoinnetwork"
)

var (
	ParamsKey = []byte("p_raficoinnetwork")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
