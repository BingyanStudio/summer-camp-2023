package encryption

import (
    "crypto/sha256"
	"encoding/hex"
    "fmt"
)

func Myencrypt(row string) string{
    data := []byte(row)
    hash := sha256.Sum256(data)
    fmt.Printf("%x\n", hash)
	res := hex.EncodeToString(hash[:])
	return res
}