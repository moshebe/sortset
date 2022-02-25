package set

import (
	"fmt"
	"sort"
) 

type Set[T any] struct {
  items []T
  key func(item T) string
  sort func(a, b T) bool
}

func New[T any](key func(item T) string, sort func(a, b T) bool) (*Set[T], error) {
	if key == nil {
		return nil, fmt.Errorf("missing uniqueness key function")
	}

	if sort == nil {
		return nil, fmt.Errorf("missing sort function")
	}

	return &Set[T]{ key: key, sort: sort}, nil
}

func (s *Set[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s *Set[T]) Sort() error {
	if s.sort == nil {
		return fmt.Errorf("missing sort function")
	}
	sort.SliceStable(s.items, func(i, j int) bool {
		return s.sort(s.items[j], s.items[i])		
	})
	return nil
}

func (s *Set[T]) Unique() []T {
	results := make([]T, 0, len(s.items))
	uniqueKeys := make(map[string]struct{})

	for _, item := range s.items {
		key := s.key(item)
		if _, found := uniqueKeys[key]; found {
			continue
		}
		uniqueKeys[key] = struct{}{}
		results = append(results, item)
	}

	return results
}