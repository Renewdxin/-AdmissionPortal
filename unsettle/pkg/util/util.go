package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodingMD5(code string) string {

	// 创建一个 MD5 的哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并写入哈希对象
	_, err := hasher.Write([]byte(code))
	if err != nil {
		return ""
	}

	// 计算哈希值并将其转换为16进制字符串
	hashInBytes := hasher.Sum(nil)
	md5String := hex.EncodeToString(hashInBytes)

	return md5String
}
