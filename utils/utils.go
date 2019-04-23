package utils

import (
	"fmt"
	"os"
)

func WriteFile( fileName string, WriteMode byte, data string ) (ok bool, err error) {
	// TODO 向目标fileName文件，写入/追加 字符串data

	var mode int

	if WriteMode == 'w' {
		mode = os.O_WRONLY|os.O_TRUNC|os.O_CREATE
	} else if WriteMode == 'a' {
		mode = os.O_RDWR|os.O_CREATE|os.O_APPEND
	} else {
		return false, nil
	}

	destFile, err := os.OpenFile(fileName, mode, 0644 )

	if destFile != nil { defer destFile.Close() }
	if err != nil {
		fmt.Println(err)
		return false, nil // 写入失败
	}

	destFile.WriteString( data )

	// 写入成功
	return true, nil
}

