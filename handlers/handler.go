package handlers

import (
	"aio/helpers"
	"aio/logging"
	"aio/modules/management"
	"aio/modules/mass"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/zenthangplus/goccm"
)

func ConverterHandler(auths []string, proxies []string, threads int) {
	c := goccm.New(threads)
	for _, token := range auths {
		c.Wait()
		// Log the entire message
		go func() {
			authToken := management.AuthTokens(token, proxies)
			if authToken != "" {
				// Append the auth token to the file
				go helpers.AppendToFile("output/converter/tokens.txt", fmt.Sprintf("%s\n", authToken))
				go helpers.AppendToFile("input/tokens.txt", fmt.Sprintf("%s\n", authToken))
				// Log the specific parts with color formatting
				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("converter ")+color.GreenString("token=")+color.WhiteString(strings.Split(authToken, ":")[0])+color.GreenString("status=")+color.WhiteString(" success"))
			} else {
				helpers.AppendToFile("output/converter/failed.txt", fmt.Sprintf("%s\n", authToken))
				logging.Log(logging.Warning, color.CyanString("module=")+color.WhiteString("converter ")+color.CyanString("token=")+color.WhiteString(token)+color.CyanString("module=")+color.WhiteString(" failed"))

			}
			c.Done()
		}()
	}
	c.WaitAllDone()
	return
}

func CheckerHandler(tokens []string, threads int) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := management.CheckTokenNew(auth_token, ct0, proxy)
			if err != nil {
				helpers.AppendToFile("output/checker/failed.txt", fmt.Sprintf("%s\n", t))
				c.Done()
				return
			}
			if status == "" {
				helpers.AppendToFile("output/checker/failed.txt", fmt.Sprintf("%s\n", t))
				c.Done()
				return
			}
			if status == "VALID" {
				helpers.AppendToFile("output/checker/valid.txt", fmt.Sprintf("%s\n", t))

			}
			if status == "UNLOCKABLE" {
				helpers.AppendToFile("output/checker/locked.txt", fmt.Sprintf("%s\n", t))
			}
			if status == "CONSENT" {
				helpers.AppendToFile("output/checker/consent.txt", fmt.Sprintf("%s\n", t))
			}
			if status == "SUSPENDED" {
				helpers.AppendToFile("output/checker/suspended.txt", fmt.Sprintf("%s\n", t))
			}
			logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("checker ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString(status))
			c.Done()
		}()

	}
	c.WaitAllDone()
	return

}

func HandleLike(tokens []string, tweet_id string, threads int) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := mass.Dolike(auth_token, ct0, proxy, tweet_id)
			if err != nil {
				c.Done()
				helpers.AppendToFile("output/like/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Error, color.GreenString("module=")+color.WhiteString("like ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))
				return
			}
			if status == "" {
				c.Done()
				helpers.AppendToFile("output/like/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("like ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			if status == "OK" {
				c.Done()
				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("like ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("SUCCESS"))

			}
			if status == "NOTOK" {
				c.Done()
				helpers.AppendToFile("output/like/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("like ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			c.Done()
		}()
	}
	c.WaitAllDone()
	return
}
func HandleRT(tokens []string, tweet_id string, threads int) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := mass.DoRetweet(auth_token, ct0, proxy, tweet_id)
			if err != nil {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Error, color.GreenString("module=")+color.WhiteString("retweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))
				c.Done()
				return
			}
			if status == "" {
				c.Done()

				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("retweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			if status == "OK" {
				c.Done()

				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("retweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("SUCCESS"))

				return
			}
			if status == "NOTOK" {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("retweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
		}()

	}
	c.WaitAllDone()
	return
}

func AiHandle(prompt string, count int, threads int) {
	c := goccm.New(threads)
	for i := 0; i < count; i++ {
		c.Wait()
		go func() {
			comment, err := mass.GenerateAIComments(prompt)
			if err != nil {
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("AiComments ")+color.GreenString(" status=")+color.WhiteString("ERROR"))
				c.Done()
				return
			}
			if comment == "" {
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("AiComments ")+color.GreenString(" status=")+color.WhiteString("EMPTY"))
				c.Done()
				return
			}
			if len(comment) < 100 {
				c.Done()
				return
			}
			logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("AiComments ")+color.GreenString(" status=")+color.WhiteString(fmt.Sprintf("%s...\n", comment[:100])))
			helpers.AppendToFile("output/comments/generated.txt", fmt.Sprintf("%s\n", comment))
			c.Done()
			return
		}()
	}
	c.WaitAllDone()
	return
}

func HandleTweets(tokens []string, tweets []string, threads int) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := mass.DoTweet(auth_token, ct0, proxy, tweets)
			if err != nil {
				c.Done()
				helpers.AppendToFile("output/tweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Error, color.GreenString("module=")+color.WhiteString("tweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))
				return
			}
			if status == "" {
				c.Done()

				helpers.AppendToFile("output/tweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("tweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			if status == "OK" {
				c.Done()

				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("tweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("SUCCESS"))

				return
			}
			if status == "NOTOK" {
				c.Done()
				helpers.AppendToFile("output/tweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("tweet ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
		}()
	}
	c.WaitAllDone()
	return
}

func HandleReply(tokens []string, tweet_id string, threads int) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := mass.DoReplyNew(auth_token, ct0, proxy, tweet_id)
			if err != nil {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Error, color.GreenString("module=")+color.WhiteString("reply")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))
				c.Done()
				return
			}
			if status == "" {
				c.Done()

				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("reply")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			if status == "OK" {
				c.Done()

				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("reply")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("SUCCESS"))

				return
			}
			if status == "NOTOK" {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("reply")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
		}()

	}
	c.WaitAllDone()
	return
}

func HandleFollow(tokens []string, threads int, screen_name string) {
	c := goccm.New(threads)
	for _, t := range tokens {
		c.Wait()
		go func() {
			if strings.Count(t, ":") != 5 {
				c.Done()
				return
			}
			auth_token := strings.Split(t, ":")[0]
			ct0 := strings.Split(t, ":")[1]
			proxy := strings.Split(t, ":")[2] + ":" + strings.Split(t, ":")[3] + ":" + strings.Split(t, ":")[4] + ":" + strings.Split(t, ":")[5]
			status, err := mass.DoFollow(auth_token, ct0, proxy, screen_name)
			if err != nil {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Error, color.GreenString("module=")+color.WhiteString("follow ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))
				c.Done()
				return
			}
			if status == "" {
				c.Done()

				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("follow ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
			if status == "OK" {
				c.Done()

				logging.Log(logging.Info, color.GreenString("module=")+color.WhiteString("follow ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("SUCCESS"))

				return
			}
			if status == "NOTOK" {
				c.Done()
				helpers.AppendToFile("output/retweet/failed.txt", fmt.Sprintf("%s\n", t))
				logging.Log(logging.Warning, color.GreenString("module=")+color.WhiteString("follow ")+color.GreenString("token=")+color.WhiteString(auth_token)+color.GreenString(" status=")+color.WhiteString("ERROR"))

				return
			}
		}()

	}
	c.WaitAllDone()
	return
}
