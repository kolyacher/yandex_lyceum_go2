/*
Напишите функцию ModifyFile(filename string, pos int, val string),

func ModifyFile(filename string, pos int, val string)

которая будет изменять содержимое файла с именем filename на значение val, начиная с позиции pos. Для перемещения по файлу используйте функцию os.Seek.
*/
package main

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	file, _ := os.OpenFile(filename, os.O_RDWR, 0644)
	defer file.Close()
	_, _ = file.Seek(int64(pos), 0)
	_, _ = file.WriteString(val)
}
