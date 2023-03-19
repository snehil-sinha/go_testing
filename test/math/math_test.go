package math_test

import (
	"math"
	"os"
	"testing"
)

func TestAbs(t *testing.T) {
	// -1 0 1
	if math.Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}
	if math.Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0)
	}
	if math.Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1)
	}
}

func TestAbsSub(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if math.Abs(1) < 0 {
			t.Error("Negative value found in abs() with", 1)
		}
	})
	t.Run("Zero", func(t *testing.T) {
		if math.Abs(0) < 0 {
			t.Error("Negative value found in abs() with", 0)
		}
	})
	t.Run("Negative", func(t *testing.T) {
		if math.Abs(-1) < 0 {
			t.Error("Negative value found in abs() with", -1)
		}
	})
}

func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) == 0 {
		t.Skip("Skipping the remaining instructions")
	}
	// ...
	t.Log("Tested with GOPATH: ", os.Getenv("GOPATH"))
}
