package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

type file struct {
	path     string
	filetype byte
}

func (b *buffer) readFileToBuf(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		l := new(line)
		l.text = []rune(scanner.Text())
		b.lines = append(b.lines, l)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (b *buffer) writeBufToFile(path string) {
	content := []byte{}
	for _, l := range b.lines {
		l.text = append(l.text, '\n')
		content = append(content, string(l.text)...)
	}
	ioutil.WriteFile(path, content, os.ModePerm)
}
