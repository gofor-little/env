package env

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Load loads and sets the environment variables from file using fileNames.
func Load(fileNames ...string) error {
	envs := map[string]string{}
	var failedEnvs bytes.Buffer

	// Range over the fileNames and parse them into a map.
	for _, fileName := range fileNames {
		fileData, err := os.ReadFile(filepath.Clean(fileName))
		if err != nil {
			return err
		}

		fileEnvs, err := parse(fileData, true)
		if err != nil {
			return fmt.Errorf("failed to parse file at %s: %w", fileName, err)
		}

		for key, value := range fileEnvs {
			envs[key] = value
		}
	}

	// Range over the map and call os.Setenv on each key-value pair.
	for key, value := range envs {
		if os.Getenv(key) != "" {
			continue
		}

		if err := os.Setenv(key, value); err != nil {
			if failedEnvs.Len() != 0 {
				failedEnvs.WriteString(" ")
			}

			failedEnvs.WriteString(key)
		}
	}

	if failedEnvs.Len() != 0 {
		return fmt.Errorf("failed to set the following environment variables: %v", failedEnvs)
	}

	return nil
}

// Write writes a key-value pair to a file that can be set to an environment variable later on with
// env.Load(). If 'setAfterWrite' is true env.Set will also be called on the key-value pair.
func Write(key, value, fileName string, setAfterWrite bool) error {
	// Read the file data.
	fileData, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return err
	}

	// Parse the file data into a map.
	fileEnvs, err := parse(fileData, false)
	if err != nil {
		return fmt.Errorf("failed to parse file at %s: %w", fileName, err)
	}
	fileEnvs[key] = value
	backupFileName := fileName + ".back"

	// Write to the backup file.
	if err := os.WriteFile(backupFileName, []byte(fmt.Sprintf("%#v", fileData)), 0600); err != nil {
		return err
	}

	// Truncate the existing file so we can write the updated list of key-value pairs.
	file, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return err
	}

	// Get keys from fileEnvs so we can sort their order before we write to file.
	var keys []string
	for key := range fileEnvs {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	ok := true

	// Write the key-value pairs to file.
	for _, k := range keys {
		_, err = file.WriteString(fmt.Sprintf("%v=%v\n", k, fileEnvs[k]))
		if err != nil {
			ok = false
		}
	}

	// If writing to file was successful and setAfterWrite is true, set the key-value pair as an
	// environment variable.
	if ok && setAfterWrite {
		if err := Set(key, value); err != nil {
			ok = false
		}
	}

	// If writing to file or setting the new environment variable failed rollback the file changes.
	if !ok {
		if err := os.Rename(backupFileName, fileName); err != nil {
			return err
		}
	} else {
		if err := os.Remove(backupFileName); err != nil {
			return err
		}
	}

	return nil
}

// Get gets an environment variable with a default backup value as the second parameter.
func Get(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// MustGet gets an environment variable and will return an error if the environment
// variable is not set or is empty.
func MustGet(key string) (string, error) {
	value := os.Getenv(key)
	if value != "" {
		return value, nil
	}

	return "", fmt.Errorf("environment variable value is not set or is empty for key: %s", key)
}

// Set is just a wrapper os.Setenv, this is useful for locally overriding
// environment variables.
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// parse takes in a byte array first, splitting it into separate lines, then
// splitting those lines into key-value pairs using the '=' character as a delimiter.
// Finally the key-value pairs are returned as a map.
func parse(data []byte, stripQuotes bool) (map[string]string, error) {
	lines := strings.Split(string(data), "\n")
	envs := map[string]string{}

	for _, line := range lines {
		// skip empty lines and comments
		if line == "" || regexp.MustCompile(`^\s*#|^\s*"#`).MatchString(line) {
			continue
		}

		lineParts := strings.SplitN(line, "=", 2)

		if len(lineParts) != 2 {
			return nil, fmt.Errorf("failed to parse line, expected 2 parts got %d", len(lineParts))
		}

		key := strings.TrimSpace(lineParts[0])
		value := strings.TrimSpace(lineParts[1])

		// Define a regular expression pattern to capture string within double quotes
		regex := regexp.MustCompile(`"(.*?)"`)
		match := regex.FindString(value)

		// Removes characters that are not enclosed within double quotes in a given string.
		// if value is not contain double quotes, remove the inline comment
		if match != "" {
			value = match
		} else {
			// Define a regular expression pattern to match and capture everything after '#'
			regex := regexp.MustCompile(`#.*`)

			// Replace the input string, removing everything after '#'
			result := regex.ReplaceAllString(value, "")

			// Trim any leading and trailing spaces
			value = strings.TrimSpace(result)

		}

		if stripQuotes && value[0] == '"' && value[len(value)-1] == '"' {
			for i := range value {
				if i > 0 {
					value = value[i:]
					break
				}
			}

			value = value[:len(value)-1]
		}

		envs[key] = value
	}

	return envs, nil
}
