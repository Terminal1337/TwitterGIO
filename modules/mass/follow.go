package mass

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func DoFollow(auth_token string, ct0 string, proxy string, screen_name string) (string, error) {
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
	body := fmt.Sprintf(`include_profile_interstitial_type=1&include_blocking=1&include_blocked_by=1&include_followed_by=1&include_want_retweets=1&include_mute_edge=1&include_can_dm=1&include_can_media_tag=1&include_ext_has_nft_avatar=1&include_ext_is_blue_verified=1&include_ext_verified_type=1&include_ext_profile_image_shape=1&skip_status=1&screen_name=%s`, screen_name)
	req, err := http.NewRequest(http.MethodPost, "https://twitter.com/i/api/1.1/friendships/create.json", strings.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return "NOTOK", err
	}
	req.Header = http.Header{
		"accept":                {"*/*"},
		"accept-encoding":       {"gzip, deflate, br"},
		"authorization":         {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
		"content-type":          {"application/x-www-form-urlencoded"},
		"cookie":                {fmt.Sprintf("auth_token=%s;ct0=%s", auth_token, ct0)},
		"origin":                {"https://twitter.com"},
		"referer":               {fmt.Sprintf("https://twitter.com/%s", screen_name)},
		"sec-ch-ua-mobile":      {"?0"},
		"sec-ch-ua-platform":    {`"Windows"`},
		"sec-fetch-dest":        {"empty"},
		"sec-fetch-mode":        {"cors"},
		"sec-fetch-site":        {"same-origin"},
		"x-csrf-token":          {ct0},
		"x-twitter-active-user": {"yes"},
		"x-twitter-auth-type":   {"OAuth2Session"},
		http.HeaderOrderKey: {
			"accept",
			"accept-encoding",
			"authorization",
			"content-type",
			"cookie",
			"origin",
			"referer",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"user-agent",
			"x-csrf-token",
			"x-twitter-active-user",
			"x-twitter-auth-type",
		},
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return "NOTOK", err
	}
	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		b, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(b))
		return "NOTOK", fmt.Errorf("Status Code Error != 200 At Like Function")
	}

	return "OK", nil
}
