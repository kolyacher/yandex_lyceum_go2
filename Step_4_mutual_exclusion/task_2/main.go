/*
Напишите потокобезопасный счётчик

	type Counter struct {
	    value int
	    mu    sync.RWMutex
	}

# Реализуйте следующий интерфейс

	type Сount interface{
	    Increment() // увеличение счётчика на единицу
	    GetValue() int // получение текущего значения
	}
*/
package main

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}
type Сount interface {
	Increment()
	GetValue() int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
