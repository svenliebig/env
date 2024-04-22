package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("Load", func(t *testing.T) {
		err := Load()

		if err != nil {
			t.Errorf("Load() error = %v, want nil", err)
			return
		}

		if got := os.Getenv("HELLO"); got != "world" {
			t.Errorf("Load() = %v, want %v", got, "world")
		}
	})
}
