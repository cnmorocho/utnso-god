package utils

import "errors"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var nullItem T
		return nullItem, errors.New("la cola no tiene elementos")
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

func (q *Queue[T]) GetFirstElement() (T, error) {
	if q.IsEmpty() {
		var nullItem T
		return nullItem, errors.New("la cola no tiene elementos")
	}
	return q.items[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
