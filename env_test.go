package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("should load the env from the .env file", func(t *testing.T) {
		err := Load()

		if err != nil {
			t.Errorf("Load() error = %v, want nil", err)
			return
		}

		if got := os.Getenv("HELLO"); got != "world" {
			t.Errorf("Load() = %v, want %v", got, "world")
		}
	})

	t.Run("should prefer the system env before the the env in the .env file", func(t *testing.T) {
		expected := "westeros"

		os.Setenv("HELLO", expected)

		err := Load()

		if err != nil {
			t.Errorf("Load() error = %v, want nil", err)
			return
		}

		if got := os.Getenv("HELLO"); got != expected {
			t.Errorf("Load() = %v, want %v", got, expected)
		}
	})
}
