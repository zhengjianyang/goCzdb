package goCzdb

import (
	"encoding/base64"
	"github.com/zhengjianyang/goCzdb/hyperHeaderDecoder"
	"io"
)

type Decryptor struct {
	keyBytes []byte
}

func NewDecryptor(key string) *Decryptor {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil
	}
	return &Decryptor{keyBytes: keyBytes}
}

func (d Decryptor) decrypt(data []byte) []byte {
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ d.keyBytes[i%len(d.keyBytes)]
	}
	return result
}

/*
GetVersion
获取czdb数据库版本

file: 数据库文件

key： 密钥
*/
func GetVersion(file io.Reader, key string) (int64, error) {
	headerBlock, err := hyperHeaderDecoder.Decrypt(file, key)
	if err != nil {
		return 0, err
	}
	return headerBlock.Version, nil
}
