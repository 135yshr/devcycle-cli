package cmd

import (
	"github.com/135yshr/devcycle-cli/internal/api"
	"github.com/135yshr/devcycle-cli/internal/config"
)

func loadToken() (*api.Token, error) {
	tokenPath, err := config.GetTokenPath()
	if err != nil {
		return nil, err
	}
	return api.LoadToken(tokenPath)
}
