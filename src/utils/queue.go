package utils

import "errors"

type Queue[T any] struct {
	storage []T
}

func (q *Queue[T]) Enqueue(element T) {
	q.storage = append(q.storage, element)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var nullElement T
		return nullElement, errors.New("la cola no tiene elementos")
	}

	element := q.storage[0]
	q.storage = q.storage[1:]

	return element, nil
}

func (q *Queue[T]) GetFirstElement() (T, error) {
	if q.IsEmpty() {
		var nullElement T
		return nullElement, errors.New("la cola no tiene elementos")
	}
	return q.storage[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.storage) == 0
}
