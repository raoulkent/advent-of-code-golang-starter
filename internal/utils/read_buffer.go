package utils

import (
	"bufio"
	"fmt"
	"os"
)

func BufferFile(path string) (*bufio.Reader, error) {
	data, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer data.Close()
	reader := bufio.NewReader(data)

	return reader, nil
}
