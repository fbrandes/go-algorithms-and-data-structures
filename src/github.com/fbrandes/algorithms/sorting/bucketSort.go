package sorting

import "github.com/fbrandes/util"

var bucketSize = 10

func BucketSort(a []int) []int {
	return bucketSort(a)
}

func bucketSort(a []int) []int {
	if len(a) == 0 {
		return a
	}

	max := util.GetMax(a)
	min := util.GetMin(a)

	numBuckets := int(max-min) / bucketSize + 1
	buckets := make([][]int, numBuckets)

	for i := 0; i < numBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range a {
		index := int(n - min) / bucketSize
		buckets[index] = append(buckets[index], n)
	}

	sorted := make([]int, 0)

	for _, bucket := range buckets {
		if len(bucket) > 0 {
			InsertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}
