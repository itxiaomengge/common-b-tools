package encryption

import (
	"crypto/md5"
	"io"
	"fmt"
)

// @Title Md5ByInt
// @Description 将字符串通过 md5 加密返回
// @Author Xiaomeng.Ge
// @Date 2022-12-29 10:05:49
//
// @Param str string
//
// @return string
func Md5ByString(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)

	if err != nil {
		return "", err
	}

	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr), nil
}
