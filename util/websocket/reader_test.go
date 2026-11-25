package websocket

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReaderURL(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		pth      string
		expected string
	}{
		{
			name:     "no path",
			baseURL:  "http://example.com",
			pth:      "/api",
			expected: "ws://example.com/api",
		},
		{
			name:     "with path",
			baseURL:  "https://example.com/garm",
			pth:      "/api",
			expected: "wss://example.com/garm/api",
		},
		{
			name:     "path with trailing slash",
			baseURL:  "https://example.com/garm/",
			pth:      "/api",
			expected: "wss://example.com/garm/api",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader, err := NewReader(context.Background(), tc.baseURL, tc.pth, "token", nil)
			require.NoError(t, err)
			require.Equal(t, tc.expected, reader.url.String())
		})
	}
}
