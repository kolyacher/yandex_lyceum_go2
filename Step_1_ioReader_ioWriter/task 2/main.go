/*
Напишите функцию

func ReadString(r io.Reader) (string, error)

которая читает данные с помощью r и возвращает их в строковом виде. Если возникает ошибка, функция должна возвращать пустую строку и ошибку, иначе — строку и nil.
*/
package main

import "io"

func ReadString(r io.Reader) (string, error) {
	buf := make([]byte, 1024)
	re, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf[:re]), nil
}
