package management

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func CheckToken(authtoken string, ct0 string, proxy string) (string, error) {

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return "", err
	}
	headers := map[string]string{
		"authority":                 "twitter.com",
		"accept":                    "*/*",
		"accept-language":           "en-US,en;q=0.9",
		"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"origin":                    "https://twitter.com",
		"referer":                   "https://twitter.com/settings/profile",
		"sec-fetch-dest":            "empty",
		"sec-fetch-mode":            "cors",
		"sec-fetch-site":            "same-origin",
		"sec-gpc":                   "1",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"x-twitter-active-user":     "yes",
		"x-twitter-auth-type":       "OAuth2Session",
		"x-twitter-client-language": "en",
	}

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	req, err := http.NewRequest("POST", "https://twitter.com/i/api/1.1/account/update_profile.json", nil)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("x-csrf-token", ct0)
	req.AddCookie(&http.Cookie{Name: "auth_token", Value: authtoken})
	req.AddCookie(&http.Cookie{Name: "ct0", Value: ct0})

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if resp.StatusCode == 200 {
		return "VALID", nil
	}
	b, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(b), "https://twitter.com/account/access") {
		return "LOCKED", nil
	}
	if strings.Contains(string(b), "/i/flow/consent_flow") {
		return "CONSENT", nil
	}
	if strings.Contains(string(b), "is suspended and") {
		return "SUSPENDED", nil
	}
	return "", nil
}
