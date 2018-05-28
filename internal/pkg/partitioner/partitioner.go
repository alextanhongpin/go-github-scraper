package partitioner

import "math"

// Partition contains the start and end index of the slices to be partitioned
type Partition struct {
	Start int
	End   int
}

// New returns a map of partitions that indicates the start and end
// index of the slice to be partition - a better way of handling generic
func New(perPartition int, arr int) (map[int]Partition, int) {
	p := float64(perPartition)
	arrLen := float64(arr)
	bucket := int(math.Ceil(arrLen / p))
	out := make(map[int]Partition, bucket)

	for i := 0; i < bucket; i++ {
		min := math.Min(float64(i+1)*p, arrLen)
		out[i] = Partition{
			Start: i * int(p),
			End:   int(min),
		}
	}
	return out, bucket
}
