package main

import "strconv"

var (
	newLine = []byte("\n")
	space   = []byte(" ")
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(k T) {
	s.m[k] = struct{}{}
}

func (s *Set[T]) Delete(k T) {
	delete(s.m, k)
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Exists(k T) bool {
	_, ok := s.m[k]
	return ok
}

func parseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func count[S ~[]E, E comparable](s S, target E) (total int) {
	for _, v := range s {
		if v == target {
			total++
		}
	}
	return
}
