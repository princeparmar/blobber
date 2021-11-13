package stats

type InfraStats struct {
	CPUs int
	NumberOfGoroutines int
	HeapAlloc uint64
	HeapSys uint64
	ActiveOnChain string
}
