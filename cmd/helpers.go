package cmd

import (
	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/135yshr/devcycle-cli/pkg/api"
)

func loadToken() (*api.Token, error) {
	tokenPath, err := config.TokenFilePath()
	if err != nil {
		return nil, err
	}
	return api.LoadToken(tokenPath)
}
