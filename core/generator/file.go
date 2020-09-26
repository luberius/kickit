package generator

import (
	"bufio"
	"os"
)

func CreateFile(content string, path string, name string) error {
	os.MkdirAll(path, os.ModePerm)
	file, err := os.Create(path + name)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}
