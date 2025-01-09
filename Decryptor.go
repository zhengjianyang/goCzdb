package goCzdb

import (
	"encoding/base64"
	"github.com/zhengjianyang/goCzdb/byteUtil"
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

headerBytes: 数据库文件 byte[:12]
*/
func GetVersion(headerBytes []byte) int64 {
	version := byteUtil.GetIntLong(headerBytes, 0)
	return version
}
