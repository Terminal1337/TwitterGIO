package management

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func CheckTokenNew(authtoken string, ct0 string, proxy string) (string, error) {

	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(15),
		tls_client.WithClientProfile(profiles.Chrome_124),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(proxy),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, "https://twitter.com/i/api/1.1/account/update_profile.json", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header = http.Header{
		"Authority":                 {"twitter.com"},
		"Accept":                    {"*/*"},
		"Accept-Language":           {"en-US,en;q=0.9"},
		"Authorization":             {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
		"Origin":                    {"https://twitter.com"},
		"Referer":                   {"https://twitter.com/settings/profile"},
		"Cookie":                    {fmt.Sprintf("auth_token=%s;ct0=%s", authtoken, ct0)},
		"Sec-Fetch-Dest":            {"empty"},
		"Sec-Fetch-Mode":            {"cors"},
		"Sec-Fetch-Site":            {"same-origin"},
		"Sec-Gpc":                   {"1"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"},
		"X-CSRF-TOKEN":              {ct0},
		"X-Twitter-Active-User":     {"yes"},
		"X-Twitter-Auth-Type":       {"OAuth2Session"},
		"X-Twitter-Client-Language": {"en"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return "VALID", nil
	}
	b, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(b), "https://twitter.com/account/access") {
		return "UNLOCKABLE", nil
	}
	if strings.Contains(string(b), "/i/flow/consent_flow") {
		return "CONSENT", nil
	}
	if strings.Contains(string(b), "is suspended and") {
		return "SUSPENDED", nil
	}
	return "", nil
}
