package tempconv_test

import (
	"testing"
	"testing/quick"

	"github.com/sergiofgonzalez/tempconv"
)

func TestCtoF(t *testing.T) {
	tests := map[string]struct {
		c         tempconv.Celsius
		expectedF tempconv.Fahrenheit
	}{
		"water freezing point": {
			c:         0,
			expectedF: 32,
		},
		"water boiling point": {
			c:         100,
			expectedF: 212,
		},
		"body temp": {
			c:         37.5,
			expectedF: 99.5,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.CtoF(testCase.c)
			if got != testCase.expectedF {
				t.Errorf("expected %g, but got %g", testCase.expectedF, got)
			}
		})
	}
}

func TestFtoC(t *testing.T) {
	tests := map[string]struct {
		f         tempconv.Fahrenheit
		expectedC tempconv.Celsius
	}{
		"water freezing point": {
			f:         32,
			expectedC: 0,
		},
		"water boiling point": {
			f:         212,
			expectedC: 100,
		},
		"body temp": {
			f:         99.5,
			expectedC: 37.5,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.FtoC(testCase.f)
			if got != testCase.expectedC {
				t.Errorf("expected %g, but got %g", testCase.expectedC, got)
			}
		})
	}
}

func TestPropertiesConversion(t *testing.T) {
	t.Run("FtoC(CtoF) should match", func(t *testing.T) {
		assertion := func(f tempconv.Fahrenheit) bool {
			// Ignore temp below 0K
			if f < -459.67 || f > 1e6 {
				return true
			}
			t.Log("testing", f)
			c := tempconv.FtoC(f)
			fFromC := tempconv.CtoF(c)
			return fFromC == f
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("Property based test failed", err)
		}
	})

	t.Run("CtoF(FtoC) should match", func(t *testing.T) {
		assertion := func(c tempconv.Celsius) bool {
			// Ignore temp below 0K
			if c < 2273.15 || c > 1e6 {
				return true
			}
			t.Log("testing", c)
			f := tempconv.CtoF(c)
			cFromF := tempconv.FtoC(f)
			return cFromF == c
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("Property based test failed", err)
		}
	})
}

func TestPrintCelsius(t *testing.T) {
	tests := map[string]struct{
		c tempconv.Celsius
		expected string
	}{
		"small positive": {
			c: 1234.567,
			expected: "1234.567°C",
		},
		"small negative": {
			c: -1234.567,
			expected: "-1234.567°C",
		},
		"positive integer": {
			c: 1234,
			expected: "1234°C",
		},
		"negative integer": {
			c: -1234,
			expected: "-1234°C",
		},
		"large positive": {
			c: 987654321.123,
			expected: "9.87654321123e+08°C",
		},
		"large negative": {
			c: -987654321.123,
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
	tests := map[string]struct{
		f tempconv.Fahrenheit
		expected string
	}{
		"small positive": {
			f: 1234.567,
			expected: "1234.567°F",
		},
		"small negative": {
			f: -1234.567,
			expected: "-1234.567°F",
		},
		"positive integer": {
			f: 1234,
			expected: "1234°F",
		},
		"negative integer": {
			f: -1234,
			expected: "-1234°F",
		},
		"large positive": {
			f: 987654321.123,
			expected: "9.87654321123e+08°F",
		},
		"large negative": {
			f: -987654321.123,
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
