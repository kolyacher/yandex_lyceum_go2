/*
Дан сервер, доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

    код 200 и значение оценки студента, если всё прошло успешно
    код 404, если студент не найден
    код 500, если у сервера проблема

Напишите функцию Average(names []string) (int, error), которая выводит среднюю успеваемость студентов с именами names. Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Average(names []string) (int, error) {
	st := "http://localhost:8082/mark?name="
	var sum int = 0
	var col int = 0
	for _, name := range names {
		st1 := st + name
		resp1, err1 := http.Get(st1)
		defer resp1.Body.Close()
		if err1 == nil {
			s1, _ := io.ReadAll(resp1.Body)
			c1, _ := strconv.Atoi(string(s1))
			sum += c1
			col += 1
		} else {
			return sum / col, err1
		}
		if resp1.StatusCode == 404 {
			return sum / col, fmt.Errorf("student %s not found", name)

		}
		if resp1.StatusCode == 500 {
			return sum / col, fmt.Errorf("server error")

		}
	}
	return sum / col, nil
}

func main() {
	var s []string = []string{"student1", "student2"}
	i, err := Average(s)
	fmt.Println(i, err)
}
