package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config

func InitGoogleAuth() {
	GoogleOauthConfig = &oauth2.Config{
		ClientID:     GetEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: GetEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  GetEnv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
		Endpoint: google.Endpoint,
	}
}
