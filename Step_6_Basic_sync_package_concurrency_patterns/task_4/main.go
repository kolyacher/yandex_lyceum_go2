/*
Дан сервер доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

	код 200 и значение оценки студента, если всё прошло успешно
	код 404, если студент не найден
	код 500, если у сервера проблема

Напишите функцию CompareList(names []string) (map[string]string, error), которая выводит карту. Её ключ — имя студента из списка name, а значение — > (оценка студента больше средней оценки студентов), < (оценка студента меньше средней оценки студентов) или = (оценка студента равна средней оценки студентов). Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

func CompareList(names []string) (map[string]string, error) {
	st := "http://localhost:8082/mark?name="
	var sum int = 0
	var col int = 0
	var sr int = 0
	sort.Slice(names, func(i, j int) bool {
		if names[i] < names[j] {
			return true
		} else {
			return false
		}
	})
	names1 := make(map[string]string)
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
			return names1, err1
		}
		if resp1.StatusCode == 404 {
			return names1, fmt.Errorf("student %s not found", name)

		}
		if resp1.StatusCode == 500 {
			return names1, fmt.Errorf("server error")

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
			names1[name] = ">"
		} else if c1 < sr {
			names1[name] = "<"
		} else {
			names1[name] = "="
		}
	}
	return names1, nil

}

func main() {
	var s []string = []string{"student2", "student1"}
	i, err := CompareList(s)
	fmt.Println(i, err)
}
