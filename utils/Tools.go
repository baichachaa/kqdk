package utils

import (
	"encoding/json"
	"math/rand"
	"time"
)

func JsonTOStructFromByte(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func StructToJsonByte(v interface{}) []byte {
	b, _ := json.MarshalIndent(v, "", "  ")
	return b
}

func StringReverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func RandomStr(n int) string {
	var l = []rune("ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	lenght := len(l)
	for i := range b {
		b[i] = l[rand.Intn(lenght)]
	}
	return string(b)
}
