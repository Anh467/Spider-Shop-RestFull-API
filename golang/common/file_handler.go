package common

import (
	"fmt"
	"io"
	"os"
)

func GetFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileContent(file *os.File) (string, error) {
	file.Seek(0, io.SeekStart)
	var content string = ""
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		content += string(buffer[:n])
	}
	return content, nil
}
