package ping_service

import (
	"testing"
)

func TestPing(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"test", "pong test"},
		{"hello", "pong hello"},
		{"", "pong "},
	}

	for _, tc := range testCases {
		output := Ping(tc.input)

		if output != tc.expected {
			t.Errorf("ping(%s) = %s; want %s", tc.input, output, tc.expected)
		}
	}
}
