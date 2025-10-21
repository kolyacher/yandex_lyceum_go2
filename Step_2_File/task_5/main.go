/*
Представьте, что программа пишет лог-файлы, где каждая строка начинается с даты формата dd.MM.YYYY. Ваша задача — написать функцию

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error)

которая вернёт строки «лога», которые созданы в указанный диапазон времени [start..end].
Например, для исходного файла:

12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info

Если start = 13.12.2022, end = 15.12.2022, функция должна вернуть:

13.12.2022 info
14.12.2022 info
15.12.2022 info

Если ни одна строка не попала в указанный диапазон, должна вернуться ошибка.
*/
package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, err := os.Open(inputFileName)
	var str []string
	if err != nil {
		return str, err
	}
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		s1 := strings.Split(s, " ")
		date, _ := time.Parse("02.01.2006", s1[0])
		if (date.After(start) || date.Equal(start)) && (end.After(date) || end.Equal(date)) {
			str = append(str, s)
		}
	}
	if len(str) == 0 {
		return str, errors.New("Slice empty!")
	}
	return str, nil
}

/*
func main() {
	var start, end time.Time
	start = time.Date(2022, 12, 13, 0, 0, 0, 0, time.UTC)
	end = time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC)
	str, err := ExtractLog("test.txt", start, end)
	if err != nil {
		fmt.Errorf("This Error %s", err)
	} else {
		fmt.Println(str)
	}
}
*/
