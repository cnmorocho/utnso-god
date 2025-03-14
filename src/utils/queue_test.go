package utils

import "testing"

func TestEnqueue(t *testing.T) {
	t.Run("Dado un elemento, cuando Queue no tiene elementos, se le agrega el elemento a la Queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		if queue.IsEmpty() {
			t.Error("la cola no tiene elementos")
		}
	})

	t.Run("Dado un elemento, cuando la Queue tiene un elemento, se le agrega el elemento y este esta en la segunda posición", func(t *testing.T) {
		first := 1
		second := 2
		queue := NewQueue[int]()
		queue.Enqueue(first)
		if queue.IsEmpty() {
			t.Fatal("la cola no tiene elementos")
		}
		queue.Enqueue(second)
		firstElement, err := queue.GetFirstElement()
		if err != nil {
			t.Fatal("hubo un error al intentar ver el primer elemento")
		}
		if firstElement == second {
			t.Errorf("se esperaba %d pero se obtuvo %d", first, second)
		}

	})
}

func TestDequeue(t *testing.T) {
	t.Run("Dado una queue, cuando esta vacia, debería lanzar un error", func(t *testing.T) {
		queue := NewQueue[int]()
		_, err := queue.Dequeue()
		if err == nil {
			t.Fatal("se esperaba un error y no se obtuvo ninguno")
		}
		if err != QueueIsEmptyError {
			t.Errorf("se esperaba el %v pero se obtuvo %v", QueueIsEmptyError, err)
		}
	})
}
