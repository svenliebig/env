package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// load reads the .env file in the current directory and sets the environment variables.
// the .env file should be in the format of KEY=VALUE.
//
// the function will only use keys that are listing in the .env file, however, if the key
// is already set in the system environment, it will be overwritten by the system environment.
// only variables that are set in the .env file, will be loaded from the system environment.
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

		key := values[0]

		if value, ok := os.LookupEnv(key); ok {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("env: error setting environment variable: %v", err)
			}

			continue
		}

		if err := os.Setenv(key, values[1]); err != nil {
			return fmt.Errorf("env: error setting environment variable: %v", err)
		}
	}

	return nil
}
