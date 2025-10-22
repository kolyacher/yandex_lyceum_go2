/*
Напишите функцию

func Write(num int)

которая записывает данные в буфер Buf []int.

# Напишите функцию

func Consume() int

которая будет забирать первое значение из буфера и возвращать его. Используйте мьютекс для синхронизации доступа к буферу.

сделайте буфер и мьютекс переменными

var (

	Buf   []int
	mutex sync.Mutex

)
*/
package main

import "sync"

var (
	Buf   []int
	mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	defer mutex.Unlock()
	Buf = append(Buf, num)
}
func Consume() int {
	mutex.Lock()
	defer mutex.Unlock()
	if len(Buf) == 0 {
		return -1
	} else {
		i := Buf[0]
		Buf = Buf[1:]
		return i
	}
}
