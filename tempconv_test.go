package tempconv_test

import (
	"testing"

	"github.com/sergiofgonzalez/tempconv"
)

func TestPrintCelsius(t *testing.T) {
	tests := map[string]struct {
		c        tempconv.Celsius
		expected string
	}{
		"small positive": {
			c:        1234.567,
			expected: "1234.567°C",
		},
		"small negative": {
			c:        -1234.567,
			expected: "-1234.567°C",
		},
		"positive integer": {
			c:        1234,
			expected: "1234°C",
		},
		"negative integer": {
			c:        -1234,
			expected: "-1234°C",
		},
		"large positive": {
			c:        987654321.123,
			expected: "9.87654321123e+08°C",
		},
		"large negative": {
			c:        -987654321.123,
			expected: "-9.87654321123e+08°C",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.c.String()
			if got != testCase.expected {
				t.Errorf("expected %q, but got %q", testCase.expected, got)
			}
		})
	}
}

func TestPrintFahrenheit(t *testing.T) {
	tests := map[string]struct {
		f        tempconv.Fahrenheit
		expected string
	}{
		"small positive": {
			f:        1234.567,
			expected: "1234.567°F",
		},
		"small negative": {
			f:        -1234.567,
			expected: "-1234.567°F",
		},
		"positive integer": {
			f:        1234,
			expected: "1234°F",
		},
		"negative integer": {
			f:        -1234,
			expected: "-1234°F",
		},
		"large positive": {
			f:        987654321.123,
			expected: "9.87654321123e+08°F",
		},
		"large negative": {
			f:        -987654321.123,
			expected: "-9.87654321123e+08°F",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.f.String()
			if got != testCase.expected {
				t.Errorf("expected %q, but got %q", testCase.expected, got)
			}
		})
	}
}
