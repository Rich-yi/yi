package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Inter(sets ...*Set) *Set {
	s.RLock()
	defer s.RUnlock()
	resultSet := New()
	for k, _ := range s.m {
		isInter := true
		for _, set := range sets {
			if !set.Has(k) {
				isInter = false
				break
			}
		}
		if isInter {
			resultSet.Add(k)
		}
	}
	return resultSet
}

func main() {
	s1 := New()
	s1.Add(1)
	s1.Add(2)

	s2 := New()
	s2.Add(2)
	s2.Add(3)

	fmt.Println(s1.Inter(s2).m)
}
