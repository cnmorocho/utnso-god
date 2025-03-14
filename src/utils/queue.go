package utils

import (
	"errors"
	"fmt"
)

var QueueIsEmptyError = errors.New("la cola no tiene elementos")

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
		return nullItem, QueueIsEmptyError
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

func (q *Queue[T]) GetFirstElement() (T, error) {
	if q.IsEmpty() {
		var nullItem T
		return nullItem, QueueIsEmptyError
	}
	return q.items[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) String() string {
	return fmt.Sprint(q.items)
}
