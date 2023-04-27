package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"math/big"
)

// 国密摘要sm3
func EncryptSm3(text string) []byte {
	h := sm3.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return sum
}

func EncryptSm3StringHex(text string) string {
	sum := EncryptSm3(text)
	return hex.EncodeToString(sum)
}

// RSA公钥加密
func EncryptRsa(text string, eStr string, nStr string) []byte {
	e, _ := new(big.Int).SetString(eStr, 16)
	n, _ := new(big.Int).SetString(nStr, 16)
	c := new(big.Int).SetBytes([]byte(text))
	sum := c.Exp(c, e, n).Bytes()
	return sum
}

func EncryptRsaStringHex(pwd string, eStr string, nStr string) string {
	sum := EncryptRsa(pwd, eStr, nStr)
	return hex.EncodeToString(sum)
}

const (
	C1C2C3 = 1
	C1C3C2 = 0
)

// 国密sm2 椭圆曲线 加密
func EncryptSm2(model int, pubKeyStr, text string) []byte {
	x, _ := new(big.Int).SetString(pubKeyStr[:64], 16)
	y, _ := new(big.Int).SetString(pubKeyStr[64:], 16)
	pub := &sm2.PublicKey{X: x, Y: y, Curve: sm2.P256Sm2()}
	cipherTxt, _ := sm2.Encrypt(pub, []byte(text), rand.Reader, model)
	return cipherTxt
}

func EncryptSm2StringHex(model int, pubKeyStr, text string) string {
	sum := EncryptSm2(model, pubKeyStr, text)
	return hex.EncodeToString(sum)
}

// EncryptSHA256 哈希摘要sha-256
func EncryptSHA256(text string) []byte {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}

func EncryptSHA256StringHex(text string) string {
	sum := EncryptSHA256(text)
	return hex.EncodeToString(sum)
}

// EncryptAesCBCStringToBase64 AES加密
func EncryptAesCBCStringToBase64(data, key string) string {
	pwdB := []byte(data)
	keyB := []byte(key)
	block, _ := aes.NewCipher(keyB)

	padnum := block.BlockSize() - len(pwdB)%block.BlockSize()
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	pwdB = append(pwdB, pad...)

	blockMode := cipher.NewCBCEncrypter(block, keyB)
	blockMode.CryptBlocks(pwdB, pwdB)
	return base64.StdEncoding.EncodeToString(pwdB)
}

// DecryptAesCBCBase64ToString AES解密
func DecryptAesCBCBase64ToString(rowData, key string) string {

	data, err := base64.StdEncoding.DecodeString(rowData)
	if err != nil {
		return ""
	}

	keyB := []byte(key)
	block, _ := aes.NewCipher(keyB)

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyB[:blockSize])

	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	length := len(crypted)
	if length == 0 {
		return ""
	}
	//获取填充的个数
	unPadding := int(crypted[length-1])
	d := crypted[:(length - unPadding)]

	return string(d)
}

// EncryptMd5ToString md5 哈希
func EncryptMd5ToString(src string) string {
	data := []byte(src)
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
