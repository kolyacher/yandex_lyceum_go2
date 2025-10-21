/*
Напишите функцию

func Contains(r io.Reader, seq []byte) (bool, error)

которая должна найти в данных первое вхождение байт seq, которые доступны через Reader r.

Если последовательность найдена, программа возвращает true, nil, иначе false, nil. Если возникает ошибка, функция должна возвращать false и ошибку.
*/

package main

import (
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	buf := make([]byte, 4096)
	var leftover []byte

	for {
		n, err := r.Read(buf)
		if n > 0 {
			data := append(leftover, buf[:n]...)
			if bytes.Contains(data, seq) {
				return true, nil
			}
			if len(data) >= len(seq)-1 {
				leftover = data[len(data)-(len(seq)-1):]
			} else {
				leftover = data
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
	}

	return false, nil
}
