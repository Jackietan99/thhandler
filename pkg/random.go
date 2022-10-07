package pkg

import (
	cryptoRand "crypto/rand"
	"math/big"
	"strings"
)

const (
	letterBytes    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes    = "0123456789"
	letterBytesLen = len(letterBytes)
	numberBytesLen = len(numberBytes)
)

/**
 * @Description: 真随机 数字
 */
func RealRandNumber(size int) string {
	var buf strings.Builder

	for i := 0; i < size; i++ {
		result, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(numberBytesLen)))
		if err != nil {
			return ""
		}

		index := int(result.Int64())
		buf.WriteString(numberBytes[index : index+1])
	}
	str := buf.String()

	return str
}
