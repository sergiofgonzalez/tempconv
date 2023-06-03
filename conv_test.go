package tempconv_test

import (
	"math"
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

func TestCtoK(t *testing.T) {
	tests := map[string]struct {
		c         tempconv.Celsius
		expectedK tempconv.Kelvin
	}{
		"water freezing point": {
			c:         0,
			expectedK: 273.15,
		},
		"water boiling point": {
			c:         100,
			expectedK: 373.15,
		},
		"body temp": {
			c:         37.5,
			expectedK: 310.65,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.CtoK(testCase.c)
			if got != testCase.expectedK {
				t.Errorf("expected %g, but got %g", testCase.expectedK, got)
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

func TestFtoK(t *testing.T) {
	tests := map[string]struct {
		f         tempconv.Fahrenheit
		expectedK tempconv.Kelvin
	}{
		"water freezing point": {
			f:         32,
			expectedK: 273.15,
		},
		"water boiling point": {
			f:         212,
			expectedK: 373.15,
		},
		"body temp": {
			f:         99.5,
			expectedK: 310.65,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.FtoK(testCase.f)
			if got != testCase.expectedK {
				t.Errorf("expected %g, but got %g", testCase.expectedK, got)
			}
		})
	}
}

func TestKtoC(t *testing.T) {
	tests := map[string]struct {
		k         tempconv.Kelvin
		expectedC tempconv.Celsius
	}{
		"water freezing point": {
			k:         273.15,
			expectedC: 0,
		},
		"water boiling point": {
			k:         373.15,
			expectedC: 100,
		},
		"body temp": {
			k:         310.65,
			expectedC: 37.5,
		},
		"absolute zero": {
			k:         0,
			expectedC: -273.15,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.KtoC(testCase.k)
			if got != testCase.expectedC {
				t.Errorf("expected %g, but got %g", testCase.expectedC, got)
			}
		})
	}
}

func TestKtoF(t *testing.T) {
	tests := map[string]struct {
		k         tempconv.Kelvin
		expectedF tempconv.Fahrenheit
	}{
		"water freezing point": {
			k:         273.15,
			expectedF: 32,
		},
		"water boiling point": {
			k:         373.15,
			expectedF: 212,
		},
		"body temp": {
			k:         310.65,
			expectedF: 99.5,
		},
		"absolute zero": {
			k:         0,
			expectedF: -459.67,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := tempconv.KtoF(testCase.k)
			if math.Abs(float64(got - testCase.expectedF)) > 0.001 {
				t.Errorf("expected %g, but got %g", testCase.expectedF, got)
			}
		})
	}
}

func TestPropertiesConversion(t *testing.T) {
	t.Run("Conversion from Fahrenheit", func(t *testing.T) {
		assertion := func(f tempconv.Fahrenheit) bool {
			// Ignore temp below 0K
			if f < -459.67 || f > 1e6 {
				return true
			}
			t.Log("testing", f)
			c := tempconv.FtoC(f)
			k := tempconv.FtoK(f)
			fFromC := tempconv.CtoF(c)
			fFromK := tempconv.KtoF(k)
			return fFromC == f && fFromK == f
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("Property based test failed", err)
		}
	})

	t.Run("Conversion from Celsius", func(t *testing.T) {
		assertion := func(c tempconv.Celsius) bool {
			// Ignore temp below 0K
			if c < -273.15 || c > 1e6 {
				return true
			}
			t.Log("testing", c)
			f := tempconv.CtoF(c)
			k := tempconv.CtoK(c)
			cFromF := tempconv.FtoC(f)
			cFromK := tempconv.KtoC(k)
			return cFromF == c && cFromK == c
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("Property based test failed", err)
		}
	})

	t.Run("Conversion from Kelvin", func(t *testing.T) {
		assertion := func(k tempconv.Kelvin) bool {
			// Ignore temp below 0K
			if k < 0 || k > 1e6 {
				return true
			}
			t.Log("testing", k)
			f := tempconv.KtoF(k)
			c := tempconv.KtoC(k)
			kFromF := tempconv.FtoK(f)
			kFromC := tempconv.CtoK(c)
			return kFromF == k && kFromC == k
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("Property based test failed", err)
		}
	})
}
