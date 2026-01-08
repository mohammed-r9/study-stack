package utils

import (
	"context"
	"fmt"
	"net/http"
	"study-stack/internal/entities/tokens/stateless"
	"study-stack/internal/shared/consts"

	"github.com/medama-io/go-useragent"
)

func SetRefreshCookie(w http.ResponseWriter, cookieValue string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(consts.TTL_REFRESH),
	})
}

func SetCsrfCookie(w http.ResponseWriter, cookieValue string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF_token",
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(consts.TTL_CSRF),
	})

}

var ua = useragent.NewParser()

// parses and returns `Browser-Name for desktop/mobile, OS` from a user-agent string
func GetDeviceNameFromUserAgent(userAgentStr string) string {
	agent := ua.Parse(userAgentStr)

	browser := agent.Browser().String()
	if browser == "" {
		browser = "Unknown Browser"
	}

	device := agent.Device().String()
	if device == "" {
		device = "Unknown Device"
	}

	os := agent.OS().String()
	if os == "" {
		os = "Unknown OS"
	}

	return fmt.Sprintf("%s for %s, %s", browser, device, os)
}

func DataFromContext(ctx context.Context) (stateless.UserClaims, bool) {
	data, ok := ctx.Value(consts.UserDataContextKey).(stateless.UserClaims)
	return data, ok
}
