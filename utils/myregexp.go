package utils

import (
	"bytes"
	"fmt"
	"regexp"
)

func regest() {

	// 以下都是做匹配

	fmt.Println(regexp.Match("H.*", []byte("Hello World!")))

	r := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(regexp.MatchReader("H.*", r))

	fmt.Println(regexp.MustCompile("H.*").MatchString("Hello World!")) // true

	// 以下是做查找
	str := "ab001234hah120210a880218end"

	// 默认的Find都是查找的byte数组，即字节数组，返回对应字符的ascii码

	data := regexp.MustCompile("\\d{6}").FindAllString(str, -1) // -1 表示全查找。

	fmt.Println(data)

	str = "hello中国hello世界和平hi好"

	data = regexp.MustCompile("[\\p{Han}]+").FindAllString(str, -1)

	fmt.Println(data)

}