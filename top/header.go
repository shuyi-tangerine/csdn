package top

import (
	"encoding/json"
	"io"
	"os"
)

type HttpHeaderType int

const (
	HttpHeaderTypeSaveArticle   = 1
	HttpHeaderTypeGetArticle    = 2
	HttpCommonHeaderPathEnvName = "COMMON_HEADERS"
)

var (
	HttpHeaderPathEnvNames = map[HttpHeaderType]string{
		HttpHeaderTypeSaveArticle: "SAVE_ARTICLE_HEADERS",
		HttpHeaderTypeGetArticle:  "GET_ARTICLE_HEADERS",
	}
)

// GetHttpHeaders 获取 http headers
// 优先从公共中获取，随后从自定义的获取覆盖填充
func GetHttpHeaders(t HttpHeaderType) (headers map[string]string, err error) {
	headers, err = GetHttpHeadersByEnvName(HttpCommonHeaderPathEnvName)
	if err != nil {
		return
	}
	tEnvName := HttpHeaderPathEnvNames[t]
	if tEnvName == "" {
		return
	}
	tHeaders, err := GetHttpHeadersByEnvName(tEnvName)
	if err != nil {
		return
	}

	if len(headers) == 0 {
		return tHeaders, nil
	}

	for k, v := range tHeaders {
		headers[k] = v
	}
	return
}

func GetHttpHeadersByEnvName(envName string) (headers map[string]string, err error) {
	headerFilePath := os.Getenv(envName)
	file, err := os.Open(headerFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	rawHeader, err := io.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(rawHeader, &headers)
	return
}
