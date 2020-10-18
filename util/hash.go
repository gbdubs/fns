package util

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func ComputeFileHash(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	hashFn := sha256.New()
	hashFn.Write([]byte(data))
	return fmt.Sprintf("%x", hashFn.Sum(nil))
}
