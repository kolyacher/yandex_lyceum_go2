/*
Напишите функцию

func Process(nums []int) chan int

которая создаёт буферизованный канал (capacity = 10) и записывает в него числа из nums без создания горутин. Функция должна вернуть созданный канал.
*/
package main

func Process(nums []int) chan int {
	ch1 := make(chan int, 10)
	for _, i := range nums {
		ch1 <- i
	}
	close(ch1)
	return ch1
}
