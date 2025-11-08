/*Напишите функцию Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error), которая должна найти первое вхождение байт seq в данных, доступных через Reader r. Если последовательность найдена, верните true, nil, иначе false, nil. В случае ошибки — false и саму ошибку. В случае отмены контекста — false и причину отмены.*/
package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	const bufSize = 4096
	buffer := make([]byte, bufSize)
	checkWindow := bytes.NewBuffer(make([]byte, 0, len(seq)))
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
		}

		n, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}
		checkWindow.Write(buffer[:n])
		if checkWindow.Len() >= len(seq) {
			windowBytes := checkWindow.Bytes()
			if bytes.Contains(windowBytes, seq) {
				return true, nil
			}
		}
		if checkWindow.Len() > len(seq)*2 {
			checkWindow.Next(checkWindow.Len() - len(seq))
		}
	}

	return false, nil
}
