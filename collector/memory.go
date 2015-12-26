package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	memory = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "memory",
		Help:      "The mem data structure holds information regarding the target system architecture of mongod and current memory use",
	}, []string{})
)

// MemStats tracks the mem stats metrics.
type MemStats struct {
	Bits              float64 `bson:"bits"`
	Resident          float64 `bson:"resident"`
	Virtual           float64 `bson:"virtual"`
	Mapped            float64 `bson:"mapped"`
	MappedWithJournal float64 `bson:"mappedWithJournal"`
}

// Export exports the data to prometheus.
func (memStats *MemStats) Export() {
	memory.WithLabelValues("resident").Set(memStats.Resident)
	memory.WithLabelValues("virtual").Set(memStats.Virtual)
	memory.WithLabelValues("mapped").Set(memStats.Mapped)
	memory.WithLabelValues("mapped_with_journal").Set(memStats.MappedWithJournal)
}
