package utils

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func ConvertStr2GBK(str string) string {
	//将utf-8编码的字符串转换为GBK编码
	ret, _ := simplifiedchinese.GBK.NewEncoder().String(str)
	return ret //如果转换失败返回空字符串
}

func ConvertGBK2Str(gbkStr string) string {
	//将GBK编码的字符串转换为utf-8编码
	ret, _ := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
	return ret //如果转换失败返回空字符串
}

func gGBKHtmlToUTF8Html(body []byte) []byte {
	html, _ := simplifiedchinese.GBK.NewDecoder().Bytes(body)
	return html
}

func GetDocumentFromGBK(body []byte) *goquery.Document {
	html := gGBKHtmlToUTF8Html(body)
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		panic(err)
	}
	return dom
}

func GetDocument(body []byte) *goquery.Document {
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	return dom
}
