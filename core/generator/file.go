package generator

import (
	"bufio"
	"os"
)

func CreateFile(content string, fullPath string) error {
	file, err := os.Create(fullPath)
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
