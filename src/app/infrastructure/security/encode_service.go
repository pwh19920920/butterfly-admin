package security

import (
	"crypto/md5"
	"fmt"
)

type EncodeServiceImpl struct {
}

func NewEncodeServiceImpl() *EncodeServiceImpl {
	return &EncodeServiceImpl{}
}

func (enc *EncodeServiceImpl) Encode(text string) string {
	return enc.MD5(enc.MD5(text))
}

// 将[]byte转成16进制
func (enc *EncodeServiceImpl) MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}
