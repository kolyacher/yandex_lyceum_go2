/*
Дан сервер, доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

	код 200 и значение оценки студента, если всё прошло успешно
	код 404, если студент не найден
	код 500, если у сервера проблема

Создайте структуру BestStudents(names []string) (string, error), которая выводит список студентов с оценками выше средней успеваемости из списка names в алфавитном порядке через запятую. Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

func BestStudents(names []string) (string, error) {
	st := "http://localhost:8082/mark?name="
	var sum int = 0
	var col int = 0
	var sr int = 0
	var str string = ""
	sort.Slice(names, func(i, j int) bool {
		if names[i] < names[j] {
			return true
		} else {
			return false
		}
	})

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
			return "", err1
		}
		if resp1.StatusCode == 404 {
			return "", fmt.Errorf("student %s not found", name)

		}
		if resp1.StatusCode == 500 {
			return "", fmt.Errorf("server error")

		}
	}
	sr = sum / col
	for _, name := range names {
		st1 := st + name
		resp1, _ := http.Get(st1)
		defer resp1.Body.Close()

		s1, _ := io.ReadAll(resp1.Body)
		c1, _ := strconv.Atoi(string(s1))
		if c1 > sr {
			if str == "" {
				str = name
			} else {
				str = str + "," + name
			}
		}
	}
	return str, nil

}

func main() {
	var s []string = []string{"student2", "student1"}
	i, err := BestStudents(s)
	fmt.Println(i, err)
}
