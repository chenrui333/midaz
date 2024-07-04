package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/LerianStudio/midaz/common/mauth"
	"github.com/coreos/go-oidc"
	"github.com/gofiber/fiber/v2"
)

const (
	GrantType    = "urn:ietf:params:oauth:grant-type:uma-ticket"
	Audience     = "midaz"
	ResponseMode = "permissions"
)

type Permissions struct {
	Scopes []string `json:"scopes"`
	Rsid   string   `json:"rsid"`
	Rsname string   `json:"rsname"`
}

type OAuthGrant struct {
	GrantType    string `json:"grant_type"`
	Audience     string `json:"audience"`
	ResponseMode string `json:"response_mode"`
}

// NewAuthnMiddleware is a function that creates a middleware to authenticate requests using JWT.
func NewAuthnMiddleware(app *fiber.App, ac *mauth.AuthClient) {

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, ac.AuthEndpoint)
	if err != nil {
		panic(err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: ac.AuthClientID, SkipClientIDCheck: true})

	app.Use(func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return Unauthorized(c, "INVALID_REQUEST", "Authorization header required")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		idToken, err := verifier.Verify(ctx, token)
		if err != nil {
			return Unauthorized(c, "INVALID_REQUEST", "Invalid token")
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			return Unauthorized(c, "INVALID_REQUEST", "Unable to parse claims")
		}

		return c.Next()
	})
}

// WithScope is a function that creates a middleware to check if the user has the required scope.
func WithScope(scope string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Get("Authorization")
		scopes, err := GetScopes(c, scope, token)

		if err != nil {
			return InternalServerError(c, "Insufficient scopes")
		}
		scopeMap := make(map[string]bool)

		for _, scope := range scopes[0].Scopes {
			scopeMap[scope] = true
		}

		if !scopeMap[scope] {
			return Forbidden(c, "Insufficient scopes")
		}

		return c.Next()
	}
}

// Dummy function to get user scopes (replace with actual implementation)
func GetScopes(c *fiber.Ctx, scope string, token string) ([]Permissions, error) {

	//Create the data
	data := url.Values{}
	data.Set("grant_type", GrantType)
	data.Set("audience", Audience)
	data.Set("response_mode", ResponseMode)

	//Create the request
	req, err := http.NewRequest("POST", os.Getenv("AUTH_ENDPOINT_TOKEN"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, InternalServerError(c, "Error creating request")
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer ", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, InternalServerError(c, "Error sending request")
	}
	defer response.Body.Close()

	// Decode the JSON response into the User struct
	// Decode the JSON response into the User struct
	var permission []Permissions
	if err := json.NewDecoder(response.Body).Decode(&permission); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	// Successfully decoded the response
	return permission, nil

}

// Fetch fetches (JWKS) JSON Web Key Set from authorization server and cache it
//
//nolint:ireturn
// func (p *JWKProvider) Fetch(ctx context.Context) (jwk.Set, error) {
// 	p.once.Do(func() {
// 		p.cache = cache.New(p.CacheDuration, p.CacheDuration)
// 	})

// 	if set, found := p.cache.Get(p.URI); found {
// 		return set.(jwk.Set), nil
// 	}

// 	set, err := jwk.Fetch(ctx, p.URI)
// 	if err != nil {
// 		return nil, err
// 	}

// 	p.cache.Set(p.URI, set, p.CacheDuration)

// 	return set, nil
// }
