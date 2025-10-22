/*
Реализуйте потокобезопасную мапу.

	type SafeMap struct {
	    m   map[string]interface{}
	    mux sync.Mutex
	}

# Для чтения элементов используйте функцию

func (s *SafeMap) Get(key string) interface{}

# Для записи элементов используйте функцию

func (s *SafeMap) Set(key string, value interface{})

# Чтобы получить новый экземпляр структуры

func NewSafeMap() *SafeMap
*/
package main

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Get(key string) interface{} {
	return s.m[key]
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}
