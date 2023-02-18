package utils

import (
	"fmt"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("文件%s不存在", fileName)
		os.Create(fileName)
	}
	return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
