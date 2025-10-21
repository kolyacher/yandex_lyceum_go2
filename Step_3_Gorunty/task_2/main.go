/*
Напишите функцию

func Receive(ch chan int) int

которая возвращает число, полученное из канала ch.
*/
package main

func Receive(ch chan int) int {
	return <-ch
}
