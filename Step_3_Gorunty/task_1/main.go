/*
Напишите программу с функцией

func Send(ch chan int, num int)

которая отправляет число num в канал ch.
*/
package main

func Send(ch chan int, num int) {
	ch <- num
}
