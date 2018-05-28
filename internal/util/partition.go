package util

import "math"

// PartitionStrings the array into buckets
func PartitionStrings(num int, arr []string) (map[int][]string, int) {
	max := float64(num)
	arrLen := float64(len(arr))
	bucket := int(math.Ceil(arrLen / max))

	out := make(map[int][]string, bucket)
	for i := 0; i < bucket; i++ {
		min := math.Min(float64(i+1)*max, arrLen)
		out[i] = arr[i*int(max) : int(min)]
	}

	return out, bucket
}
