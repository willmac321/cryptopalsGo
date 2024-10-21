package utils

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToByte(hexstr string) (int, []byte) {
	src := []byte(hexstr)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return n, dst
}

func ByteToHexString(b []byte) string {
	return hex.EncodeToString(b)
}

func ByteToBase64(bytes []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func Base64ToByte(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func Base64ByteToByte(b []byte) ([]byte, error) {
	n := base64.StdEncoding.DecodedLen(len(b))
	dst := make([]byte, n)
	base64.StdEncoding.Decode(dst, b)
	return dst, nil
}

func XOR(str []byte, compare []byte) []byte {
	res := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		res[i] = str[i] ^ compare[i%len(compare)]
	}

	return res
}

func ByteArrArrToByteArr(rep [][]byte) []byte {
	rv := []byte{}
	for _, r := range rep {
		rv = append(rv, r...)
	}
	return rv
}

func ByteArrArrToString(baa [][]byte) string {
	rv := ""
	for _, v := range baa {
		rv += string(v)
	}
	return rv
}
