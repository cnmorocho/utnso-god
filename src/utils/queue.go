package utils

import "errors"

type Queue[T any] struct {
	Items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{Items: []T{}}
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var nullItem T
		return nullItem, errors.New("la cola no tiene elementos")
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item, nil
}

func (q *Queue[T]) GetFirstElement() (T, error) {
	if q.IsEmpty() {
		var nullItem T
		return nullItem, errors.New("la cola no tiene elementos")
	}
	return q.Items[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Items) == 0
}
