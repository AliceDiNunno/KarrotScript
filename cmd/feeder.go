package main

import (
	"io"
	"os"
)

type LocalFileFeeder struct {
}

func (feeder LocalFileFeeder) ScriptRequired(filename string) ([]byte, error) {
	println("Filename:", filename)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		result = append(result, buf[:n]...)
	}
	return result, nil
}

func NewLocalFileFeeder() *LocalFileFeeder {
	return &LocalFileFeeder{}
}
