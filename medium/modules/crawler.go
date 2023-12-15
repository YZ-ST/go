package crawler

import (
	"io/ioutil"
	"net/http"
)



// 發送 HTTP 請求
func Request(url string) (*http.Response, error) {
	// 建立一個 http.Client{}
	client := &http.Client{}

	// 建立一個 http.Request{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 設置 User-Agent 標頭
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko)")


	// 發送請求
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 解析 HTTP 回應
func ParseResponse(res *http.Response) (string, error) {
	// 確保 res.Body 在函式結束時被關閉
	defer res.Body.Close()

	// 讀取回應內容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
