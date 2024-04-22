package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load reads the .env file in the current directory and sets the environment variables.
// The .env file should be in the format of KEY=VALUE.
func Load() error {
	dir, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("env: error getting current directory: %v", err)
	}

	p := fmt.Sprintf("%s/.env", dir)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return fmt.Errorf("env: .env file not found in %s", dir)
	}

	file, err := os.Open(p)

	if err != nil {
		return fmt.Errorf("env: error opening .env file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "=")

		if len(values) == 0 {
			continue
		}

		if len(values) != 2 {
			return fmt.Errorf("env: invalid .env file format")
		}

		if err := os.Setenv(values[0], values[1]); err != nil {
			return fmt.Errorf("env: error setting environment variable: %v", err)
		}
	}

	return nil
}
