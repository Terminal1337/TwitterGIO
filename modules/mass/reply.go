package mass

import (
	"fmt"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/google/uuid"
)

var ()

func DoReplyNew(auth_token string, ct0 string, proxy string, tweet_id string, message string) (string, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_124),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return "", err
	}
	body := fmt.Sprintf(`{
  "variables": {
    "tweet_text": "%s",
    "reply": {
      "in_reply_to_tweet_id": "%s",
      "exclude_reply_user_ids": []
    },
    "dark_request": false,
    "media": {
      "media_entities": [],
      "possibly_sensitive": false
    },
    "semantic_annotation_ids": []
  },
  "features": {
    "communities_web_enable_tweet_community_results_fetch": true,
    "c9s_tweet_anatomy_moderator_badge_enabled": true,
    "tweetypie_unmention_optimization_enabled": true,
    "responsive_web_edit_tweet_api_enabled": true,
    "graphql_is_translatable_rweb_tweet_is_translatable_enabled": true,
    "view_counts_everywhere_api_enabled": true,
    "longform_notetweets_consumption_enabled": true,
    "responsive_web_twitter_article_tweet_consumption_enabled": true,
    "tweet_awards_web_tipping_enabled": false,
    "creator_subscriptions_quote_tweet_preview_enabled": false,
    "longform_notetweets_rich_text_read_enabled": true,
    "longform_notetweets_inline_media_enabled": true,
    "articles_preview_enabled": true,
    "rweb_video_timestamps_enabled": true,
    "rweb_tipjar_consumption_enabled": true,
    "responsive_web_graphql_exclude_directive_enabled": true,
    "verified_phone_label_enabled": false,
    "freedom_of_speech_not_reach_fetch_enabled": true,
    "standardized_nudges_misinfo": true,
    "tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
    "responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
    "responsive_web_graphql_timeline_navigation_enabled": true,
    "responsive_web_enhance_cards_enabled": false
  },
  "queryId": "oB-5XsHNAbjvARJEc8CZFw"
}`, message, tweet_id)
	req, err := http.NewRequest(http.MethodPost, "https://x.com/i/api/graphql/oB-5XsHNAbjvARJEc8CZFw/CreateTweet", strings.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", err
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
		return "", err
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("Status Code Error != 200 At Retweet Function")
	}

	return "OK", nil
}
