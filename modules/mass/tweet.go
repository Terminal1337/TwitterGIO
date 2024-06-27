package mass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func DoTweet(auth_token string, ct0 string, proxy string, tweet_msg []string) (string, error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		// fmt.Println(err.Error())
		return "", err
	}
	type RequestPayload struct {
		Variables struct {
			TweetID string `json:"tweet_id"`
		} `json:"variables"`
		QueryID string `json:"queryId"`
	}
	type RequestBodyT struct {
		Variables struct {
			TweetID     string `json:"tweet_id"`
			DarkRequest bool   `json:"dark_request"`
		} `json:"variables"`
		QueryID string `json:"queryId"`
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

	payload := map[string]interface{}{
		"variables": map[string]interface{}{
			"tweet_text":   tweet_msg[rand.Intn(len(tweet_msg))],
			"dark_request": false,
			"media": map[string]interface{}{
				"media_entities":     []interface{}{},
				"possibly_sensitive": false,
			},
			"semantic_annotation_ids": []interface{}{},
		},
		"features": map[string]interface{}{
			"c9s_tweet_anatomy_moderator_badge_enabled":                               true,
			"tweetypie_unmention_optimization_enabled":                                true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"responsive_web_twitter_article_tweet_consumption_enabled":                false,
			"tweet_awards_web_tipping_enabled":                                        false,
			"longform_notetweets_rich_text_read_enabled":                              true,
			"longform_notetweets_inline_media_enabled":                                true,
			"rweb_video_timestamps_enabled":                                           true,
			"responsive_web_graphql_exclude_directive_enabled":                        true,
			"verified_phone_label_enabled":                                            false,
			"freedom_of_speech_not_reach_fetch_enabled":                               true,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
			"responsive_web_media_download_video_enabled":                             false,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled":       false,
			"responsive_web_graphql_timeline_navigation_enabled":                      true,
			"responsive_web_enhance_cards_enabled":                                    false,
		},
		"queryId": "bDE2rBtZb3uyrczSZ_pI9g",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", nil
	}

	req, _ := http.NewRequest("POST", "https://twitter.com/i/api/graphql/bDE2rBtZb3uyrczSZ_pI9g/CreateTweet", bytes.NewBuffer([]byte(jsonData)))
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-csrf-token", ct0)
	req.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})
	req.AddCookie(&http.Cookie{Name: "ct0", Value: ct0})
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {

		return "OK", nil
	} else {
		return "NOTOK", nil
	}

}
