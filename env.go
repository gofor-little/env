package env

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Load loads and sets the environment variables from file using filenames.
func Load(filenames ...string) error {
	envs := map[string]string{}
	var failedEnvs bytes.Buffer

	for _, filename := range filenames {
		fileData, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		fileEnvs := parse(fileData)

		for key, value := range fileEnvs {
			envs[key] = value
		}
	}

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

// Write writes a key-value pair to a file that can be set to an environment
// variable later on with env.Load().
func Write(key, value, fileName string) error {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileEnvs := parse(fileData)
	fileEnvs[key] = value
	backupFileName := fileName + ".back"

	sourceFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = sourceFile.Close() }()

	destinationFile, err := os.Create(backupFileName)
	if err != nil {
		return err
	}
	defer func() { _ = destinationFile.Close() }()

	if _, err := io.Copy(sourceFile, destinationFile); err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	var keys []string
	for key := range fileEnvs {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	writeFailed := false

	for _, k := range keys {
		_, err = file.WriteString(fmt.Sprintf("%v=%v\n", k, fileEnvs[k]))
		if err != nil {
			writeFailed = true
		}
	}

	_ = destinationFile.Close()

	if writeFailed {
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
func Get(key, def string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return def
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
func parse(data []byte) map[string]string {
	lines := strings.Split(string(data), "\n")
	envs := map[string]string{}

	for _, line := range lines {
		lineParts := strings.Split(line, "=")

		if len(lineParts) != 2 {
			continue
		}

		envs[strings.TrimSpace(lineParts[0])] = strings.TrimSpace(lineParts[1])
	}

	return envs
}
