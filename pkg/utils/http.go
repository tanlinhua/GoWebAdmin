package utils

import "github.com/go-resty/resty/v2"

// HTTP GET 请求
func HttpGet(url string, params map[string]string) (ok bool, content string) {
	client := resty.New()
	resp, err := client.R().SetQueryParams(params).Get(url)

	if err != nil || resp.StatusCode() != 200 {
		return false, err.Error()
	}

	return true, string(resp.Body())
}

// HTTP POST 请求
func HttpPost(url string, body string) (ok bool, content string) {
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).Post(url)

	if err != nil || resp.StatusCode() != 200 {
		return false, err.Error()
	}

	return true, string(resp.Body())
}
