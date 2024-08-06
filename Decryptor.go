package goCzdb

import "encoding/base64"

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
