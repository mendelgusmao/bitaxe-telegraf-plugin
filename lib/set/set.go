package set

import (
	"maps"
	"slices"
)

type set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) set[T] {
	s := make(set[T])

	for _, v := range items {
		s.Add(v)
	}

	return s
}

func (s set[T]) Add(v T) {
	if _, ok := s[v]; ok {
		return
	}

	s[v] = struct{}{}
}

func (s set[T]) Values() []T {
	return slices.Collect(
		maps.Keys[set[T]](s),
	)
}
