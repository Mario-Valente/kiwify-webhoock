package security

import (
	"testing"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
)

func TestValidadeToken(t *testing.T) {
	config := config.NewConfig()

	originalAuthSecret := config.AuthSecret
	defer func() { config.AuthSecret = originalAuthSecret }()

	tests := []struct {
		name       string
		authSecret string
		token      string
		wantValid  bool
		wantErr    bool
	}{
		{
			name:       "Valid token",
			authSecret: "secret123",
			token:      "secret123",
			wantValid:  true,
			wantErr:    false,
		},
		{
			name:       "Invalid token",
			authSecret: "secret123",
			token:      "wrongtoken",
			wantValid:  false,
			wantErr:    true,
		},
		{
			name:       "Empty token",
			authSecret: "secret123",
			token:      "",
			wantValid:  false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.AuthSecret = tt.authSecret

			gotValid, err := ValidadeToken(tt.token)
			if gotValid != tt.wantValid {
				t.Errorf("ValidadeToken() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidadeToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
