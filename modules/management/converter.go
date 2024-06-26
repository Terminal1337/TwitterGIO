package management

import (
	"aio/helpers"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func AuthTokens(accessCookie string, proxies []string) string {

	ct0, proxy, err := helpers.CsrfNew(accessCookie, proxies)
	if err == nil && ct0 != "" {
		accessCookie = fmt.Sprintf("%s:%s:%s", accessCookie, ct0, proxy)
		return accessCookie
	}
	return ""
}
func getCookieValue(cookies []*http.Cookie, cookieName string) string {
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			return cookie.Value
		}
	}
	return ""
}
func SubprocessSigner(oauth_token string, oauth_secret string) string {
	cmd := exec.Command("python", "modules/management/signer.py", oauth_token, oauth_secret)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	outputString := strings.TrimSpace(string(output))

	return outputString
}
