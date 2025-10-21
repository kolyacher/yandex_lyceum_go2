/*
Напишите функцию

func LineByNum(inputFilename string, lineNum int) string

которая получает в качестве параметров имя файла и номер строки, а возвращает текст строки по её порядковому номеру в файле. Если строки с указанным номером найти не удаётся, верните пустую строку.
*/
package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	file, _ := os.Open(inputFilename)
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		if currentLine == lineNum {
			return scanner.Text()
		}
		currentLine++
	}
	return ""
}
