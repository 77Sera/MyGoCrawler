package utils

import (
	"fmt"
	"os"
)

func WriteFile( fileName string, data string ) (ok bool, err error) {
	// TODO 向目标fileName文件，写入字符串data

	destFile, err := os.Create(fileName)

	if destFile != nil { defer destFile.Close() }
	if err != nil {
		fmt.Println(err)
		return false, nil // 写入失败
	}

	destFile.WriteString( data )

	// 写入成功
	return true, nil
}

func AppendFile( fileName string, data string ) (ok bool, err error) {
	// TODO 向目标fileName文件，追加写入字符串data

	destFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) // 追加写入文件

	if destFile != nil { defer destFile.Close() }
	if err != nil {
		fmt.Println(err)
		return false, nil // 写入失败
	}

	destFile.WriteString( data )

	// 写入成功
	return true, nil
}
