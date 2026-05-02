package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expected    string
		expectError bool
	}{
		{
			name:        "valid api key",
			header:      "ApiKey my-api-key-123",
			expected:    "my-api-key-123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			header:      "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "wrong scheme - Bearer instead of ApiKey",
			header:      "Bearer my-api-key-123",
			expected:    "",
			expectError: true,
		},
		{
			name:        "malformed - no key after ApiKey",
			header:      "ApiKey",
			expected:    "",
			expectError: true,
		},
		{
			name:        "malformed - only spaces",
			header:      "   ",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(headers)

			if tt.expectError && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
