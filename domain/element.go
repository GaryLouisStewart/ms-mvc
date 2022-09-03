package domain

type Element struct {
	Id            uint64  `json:"id"`
	Name          string  `json:"name"`
	AtomicMass    float64 `json:"atomic_mass"`
	MeltingPoint  float64 `json:"melting_point"`
	BoilingPoint  float64 `json:"boiling_point"`
	DiscoveryDate uint64  `json:"discovery_date"`
}
