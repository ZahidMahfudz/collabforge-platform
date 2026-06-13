package service

import (
	"context"
	"encoding/json"

	"github.com/zahidmahfudz/collabforge-platform/config"
)

type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type GoogleAuthService struct {
}

func NewGoogleAuthService() *GoogleAuthService {
	return &GoogleAuthService{}
}

func (s *GoogleAuthService) GetLoginURL() string {
	return config.GoogleOauthConfig.AuthCodeURL(
		"random-state",
	)
}

func (s *GoogleAuthService) GetuserByCode(code string) (*GoogleUserInfo, error) {
	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := config.GoogleOauthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var user GoogleUserInfo

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
