/*
Дан сервер, доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

	код 200 и значение оценки студента, если всё прошло успешно
	код 404, если студент не найден
	код 500, если у сервера проблема

Напишите функцию Compare(name1, name2 string) (string, error), которая сравнивает оценки двух студентов с именами name1 и name2 и выводит > (оценка студента 1 больше оценки студента 2), < (оценка студента 1 меньше оценки студента 2) или = (оценка студента 1 равна оценке студента 2). Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Compare(name1, name2 string) (string, error) {
	st := "http://localhost:8082/mark?name="
	s1 := st + name1
	s2 := st + name2
	resp1, err1 := http.Get(s1)
	defer resp1.Body.Close()
	resp2, err2 := http.Get(s2)
	defer resp2.Body.Close()
	if err1 == nil && err2 == nil {
		if resp1.StatusCode == 200 && resp2.StatusCode == 200 {
			s1, _ := io.ReadAll(resp1.Body)
			s2, _ := io.ReadAll(resp2.Body)
			c1, _ := strconv.Atoi(string(s1))
			c2, _ := strconv.Atoi(string(s2))
			if c1 > c2 {
				return ">", nil
			} else if c1 < c2 {
				return "<", nil
			} else {
				return "=", nil
			}
		}
		if resp1.StatusCode == 404 || resp2.StatusCode == 404 {
			if resp1.StatusCode == 404 {
				return "", fmt.Errorf("student %s not found", name1)
			} else {
				return "", fmt.Errorf("student %s not found", name2)
			}
		}
		if resp1.StatusCode == 500 || resp2.StatusCode == 500 {
			if resp1.StatusCode == 500 {
				return "", fmt.Errorf("server error")
			} else {
				return "", fmt.Errorf("server error")
			}
		}
	}
	return "", err2
}

func main() {
	s, err := Compare("student0", "student2")
	fmt.Println(s, err)
}
