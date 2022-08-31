package domain

type Element struct {
	Id            uint64
	Name          string
	AtomicMass    float64
	MeltingPoint  float64
	BoilingPoint  float64
	DiscoveryDate uint64
}
