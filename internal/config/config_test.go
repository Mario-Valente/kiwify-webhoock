package config

import (
	"os"
	"testing"
)

func TestCreateURIMongoDB(t *testing.T) {
	tests := []struct {
		name        string
		env         string
		mongoURL    string
		key         string
		secret      string
		expectedURI string
	}{
		{
			name:        "Development environment",
			env:         "development",
			mongoURL:    "mongodb://localhost:27017",
			key:         "admin",
			secret:      "password",
			expectedURI: "mongodb://admin:password@localhost:27017",
		},
		{
			name:        "Production environment",
			env:         "production",
			mongoURL:    "mongodb://prod.mongo.com:27017",
			key:         "prodAdmin",
			secret:      "prodSecret",
			expectedURI: "mongodb://prodAdmin:prodSecret@mongodb://prod.mongo.com:27017",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			os.Setenv("ENV", tt.env)
			os.Setenv("MONGO_URL", tt.mongoURL)
			os.Setenv("MONGO_KEY", tt.key)
			os.Setenv("MONGO_SECRET", tt.secret)

			// Create a new config
			config := NewConfig()

			// Call CreateURIMongoDB
			result := config.CreateURIMongoDB()

			// Assert the result
			if result != tt.expectedURI {
				t.Errorf("CreateURIMongoDB() = %v, want %v", result, tt.expectedURI)
			}

			// Clean up environment variables
			os.Unsetenv("ENV")
			os.Unsetenv("MONGO_URL")
			os.Unsetenv("MONGO_KEY")
			os.Unsetenv("MONGO_SECRET")
		})
	}
}
