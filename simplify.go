package simplify

import (
	"container/heap"
	"math"
)

// Point stores coordinates.
type Point struct {
	X, Y float64 // Regular coordinates.
	Z    float64 // Z-axis is used to store importance of given point in the line.
}

// area calculates area of a triangle defined by given points.
func area(a, b, c Point) float64 {
	return math.Abs((a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2)
}

// prepareHeap prepares triangles and stores them in array that will be used as
// a basis for heap.
func prepareTriangles(points []Point) TriangleHeap {
	var triangles TriangleHeap
	for i := 1; i < len(points)-1; i++ {
		// Compute the effective area of each point.
		points[i].Z = area(points[i-1], points[i], points[i+1])
		t := &Triangle{
			heap: i - 1,
			area: points[i].Z,
			a:    i - 1, b: i, c: i + 1,
		}

		// Filter out all points with zero area.
		if points[i].Z == 0 {
			continue
		}

		triangles = append(triangles, t)
	}

	return triangles
}

// connectPoints adds connections between points before they are reordered by
// heap initialization.
func connectPoints(h TriangleHeap) {
	for i := 1; i < len(h)-1; i++ {
		h[i].prev = h[i-1]
		h[i].next = h[i+1]
	}
	h[0].next = h[1]
	h[len(h)-1].prev = h[len(h)-2]

}

// Simplify provides a way to simplify array of points using Visvalingam’s
// algorithm.
func Simplify(points []Point) {
	if len(points) <= 2 {
		return
	}

	triangles := prepareTriangles(points)

	if triangles.Len() == 1 {
		triangles[0].area = area(points[0], points[1], points[2])
		points = triangles.toPointArray(points)
		return
	}

	// Initialize previous and next connections.
	connectPoints(triangles)

	heap.Init(&triangles)
	maxArea := 0.0

	// Repeat until the original line consists of only 2 points, namely the start
	// and end points which are ommited in the heap.
	for triangles.Len() > 0 {
		// Find point with the least effective area and call it the current point.
		t := triangles.Pop().(*Triangle)

		// If its calculated area is less than that of the last point to be
		// eliminated, use the latter's area instead
		// (this ensures that the current point cannot be eliminated without
		// eliminating previously eliminated points).
		if points[t.b].Z < maxArea {
			points[t.b].Z = maxArea
		} else {
			maxArea = points[t.b].Z
		}

		// Recompute the effective area of the two adjoining points.
		if t.prev != nil {
			t.prev.next = t.next
			t.prev.c = t.c
			triangles.update(t.prev, area(points[t.a], points[t.b], points[t.c]))
		} else {
			points[t.a].Z = points[t.b].Z
		}

		if t.next != nil {
			t.next.prev = t.prev
			t.next.a = t.a
			triangles.update(t.next, area(points[t.a], points[t.b], points[t.c]))
		} else {
			points[t.c].Z = points[t.b].Z
		}
	}
}
