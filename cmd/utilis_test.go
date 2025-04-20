package cmd

import (
	"testing"

	"github.com/Mario-Valente/kiwify-webhoock/internal/models"
)

func TestFormatAbandoned(t *testing.T) {
	tests := []struct {
		name     string
		input    models.Abandoned
		index    int
		expected string
	}{
		{
			name: "Valid input",
			input: models.Abandoned{
				Name:    "John Doe",
				Email:   "john.doe@example.com",
				Phone:   "+123456789",
				Country: "USA",
			},
			index:    1,
			expected: "1. 🧾 Nome: John Doe\n📧 Email: john.doe@example.com\n💰 Phone: +123456789 \n🌍 Pais: USA\n\n",
		},
		{
			name: "Empty fields",
			input: models.Abandoned{
				Name:    "",
				Email:   "",
				Phone:   "",
				Country: "",
			},
			index:    2,
			expected: "2. 🧾 Nome: \n📧 Email: \n💰 Phone:  \n🌍 Pais: \n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatAbandoned(tt.input, tt.index)
			if result != tt.expected {
				t.Errorf("formatAbandoned() = %v, want %v", result, tt.expected)
			}
		})
	}
}
