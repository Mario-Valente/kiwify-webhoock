package security

import (
	"fmt"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
)

func ValidadeToken(token string) (bool, error) {
	config := config.NewConfig()

	if config.AuthSecret == "" {
		return false, fmt.Errorf("auth secret is not set")
	}

	if token == "" {
		return false, fmt.Errorf("token is empty")
	}

	if token != config.AuthSecret {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil

}
