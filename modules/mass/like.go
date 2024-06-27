package mass

import (
	"fmt"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/google/uuid"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func Dolike(auth_token string, ct0 string, proxy string, tweet_id string) (string, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_124),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(proxy),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return "NOTOK", err
	}
	body := fmt.Sprintf(`{
  "variables": {
    "tweet_id": "%s"
  },
  "queryId": "lI07N6Otwv1PhnEgXILM7A"
}`, tweet_id)
	req, err := http.NewRequest(http.MethodPost, "https://x.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet", strings.NewReader(body))
	if err != nil {
		log.Println(err)
		return "NOTOK", err
	}
	req.Header = http.Header{
		"Accept":          {"*/*"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Authorization":   {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
		// "Connection":                {"keep-alive"},
		"Content-Type":              {"application/json"},
		"Cookie":                    {fmt.Sprintf("auth_token=%s;ct0=%s", auth_token, ct0)},
		"Host":                      {"x.com"},
		"Origin":                    {"https://x.com"},
		"Referer":                   {"https://x.com/home"},
		"Sec-Ch-Ua":                 {"\"Chromium\";v=\"124\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"124\""},
		"Sec-Ch-Ua-Mobile":          {"?0"},
		"Sec-Ch-Ua-Platform":        {"\"Windows\""},
		"Sec-Fetch-Dest":            {"empty"},
		"Sec-Fetch-Mode":            {"cors"},
		"Sec-Fetch-Site":            {"same-origin"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"},
		"X-Client-Uuid":             {uuid.New().String()},
		"X-Csrf-Token":              {ct0},
		"X-Twitter-Active-User":     {"yes"},
		"X-Twitter-Auth-Type":       {"OAuth2Session"},
		"X-Twitter-Client-Language": {""},
	}

	response, err := client.Do(req)
	if err != nil {
		return "NOTOK", err
	}
	if response.StatusCode != 200 {
		return "NOTOK", fmt.Errorf("Status Code Error != 200 At Like Function")
	}

	return "OK", nil
}
