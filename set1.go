package cryptopals

import (
	"log"
	"encoding/hex"
	"encoding/base64"
)

func HexToBase64(str string) string {
	dec, err := hex.DecodeString(str)
	if (err != nil) {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString([]byte(dec))
}
