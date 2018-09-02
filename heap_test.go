package simplify

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	triangles := make(TriangleHeap, 4)
	triangles[0] = &Triangle{heap: 0, area: 5}
	triangles[1] = &Triangle{heap: 1, area: 60}
	triangles[2] = &Triangle{heap: 2, area: 15}
	triangles[3] = &Triangle{heap: 3, area: 120}

	heap.Init(&triangles)

	t.Run("pop", func(t *testing.T) {
		triangle := triangles.Pop().(*Triangle)
		expected := 5.0

		if triangle.area != expected {
			t.Errorf(
				"Popped triangle has area %f instead of %f",
				triangle.area,
				expected,
			)
		}
	})

	t.Run("push", func(t *testing.T) {
		triangle := &Triangle{area: 10}
		triangles.Push(triangle)
		index := triangles.Len() - 1

		if triangles[index] != triangle {
			t.Errorf(
				"Heap peak should be equal to %p but there is %p",
				triangle,
				triangles[index],
			)
		}
	})

	t.Run("update", func(t *testing.T) {
		triangle := triangles[0]
		expected := 3.0
		triangles.update(triangle, expected)
		current := triangles[triangle.heap].area

		if current != expected {
			t.Errorf(
				"Heap element was not updated correctly. Current: %f. Expected: %f.",
				current,
				expected,
			)
		}
	})
}
