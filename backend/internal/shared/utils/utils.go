package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"study-stack/internal/entities/tokens/stateless"
	"study-stack/internal/shared/consts"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/medama-io/go-useragent"
)

func SetRefreshCookie(c *fiber.Ctx, cookieValue string) {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    cookieValue,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		MaxAge:   int(consts.TTL_REFRESH),
	})
}

func SetCsrfCookie(c *fiber.Ctx, cookieValue string) {
	c.Cookie(&fiber.Cookie{
		Name:     "CSRF_token",
		Value:    cookieValue,
		Path:     "/",
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Strict",
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

// extracts user claims from Fiber locals
func DataFromLocals(c *fiber.Ctx) (stateless.UserClaims, bool) {
	data, ok := c.Locals(consts.UserDataContextKey).(stateless.UserClaims)
	return data, ok
}

// takes a string uuid and return a pointer to `uuid.UUID`
func ParseOptionalUUID(s string) (*uuid.UUID, error) {
	if s == "" {
		return nil, nil
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func GenerateRandomBase64(nBytes int) (string, error) {
	if nBytes <= 0 {
		return "", fmt.Errorf("number of bytes must be > 0")
	}

	b := make([]byte, nBytes)

	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode to Base64
	return base64.RawURLEncoding.EncodeToString(b), nil
}
