package simplify

import "container/heap"

// Triangle is combination of vertex index with corresponding effective area.
type Triangle struct {
	a, b, c    int       // Points composing the triangle currently.
	heap       int       // Index in the heap.
	area       float64   // Area associated with b point.
	prev, next *Triangle // Pointers to other remaining points.
}

// TriangleHeap is a min-heap of triangles.
type TriangleHeap []*Triangle

func (h TriangleHeap) Len() int           { return len(h) }
func (h TriangleHeap) Less(i, j int) bool { return h[i].area > h[j].area }
func (h TriangleHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heap = i
	h[j].heap = j
}

// Push is used to store elements in the heap.
func (h *TriangleHeap) Push(x interface{}) {
	t := x.(*Triangle)
	t.heap = len(*h)
	*h = append(*h, t)
}

// Pop removes min element from the heap.
func (h *TriangleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	t := old[n-1]
	t.heap = -1
	*h = old[0 : n-1]
	return t
}

// update modifies the priority and value of an Item in the queue.
func (h *TriangleHeap) update(t *Triangle, area float64) {
	t.area = area
	heap.Fix(h, t.heap)
}
