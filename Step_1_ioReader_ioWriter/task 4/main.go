/*
Напишите функцию

func Copy(r io.Reader, w io.Writer, n uint) error

которая копирует n байт из r в w.

Если количество байт, доступных для чтения, меньше n, функция должна копировать все данные. Если возникает ошибка — возвращать её.
*/

package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)
	readBytes, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}

	_, err = w.Write(buf[:readBytes])
	if err != nil {
		return err
	}

	return nil
}
