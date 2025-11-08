package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeSlice struct {
	results []int
	mx      *sync.Mutex
}

func New() *SafeSlice {
	return &SafeSlice{
		mx:      &sync.Mutex{},
		results: []int{},
	}
}

// добавляем к слайсу элемент item
func (s *SafeSlice) Append(item int) {
	// вызван Lock, поэтому только одна горутина за раз может получить доступ к слайсу
	s.mx.Lock()
	defer s.mx.Unlock()
	s.results = append(s.results, random())
}

// получаем элемент слайса по индексу
func (s *SafeSlice) Get(index int) int {
	// вызван Lock, поэтому только одна горутина за раз может получить доступ к слайсу
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.results[index]
}

// функция генерирует случайное число в интервале [0, 100)
func random() int {
	const max int = 100
	return rand.Intn(max)
}

func main() {
	safeSlice := New()
	const size int = 10
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func() {
			safeSlice.Append(random())
		}()
	}
	time.Sleep(time.Second)

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(safeSlice.Get(i))
	}
}
