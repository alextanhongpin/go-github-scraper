package heapsort

import "github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

// Sort applies the heapsort sorting algorithm
func Sort(arr []schema.User) {
	n := len(arr)
	for i := n; i > -1; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func heapify(arr []schema.User, n, i int) {
	largest := i
	l := i*2 + 1
	r := i*2 + 2

	if l < n && arr[i].Score < arr[l].Score {
		largest = l
	}

	if r < n && arr[largest].Score < arr[r].Score {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
