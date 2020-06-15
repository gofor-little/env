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
			continue
		}

		fileEnvs := parse(fileData)

		for key, value := range fileEnvs {
			envs[key] = value
		}
	}

	for key, value := range envs {
		if os.Getenv(key) != "" {
			return err
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

// Write writes a key value pair to a file that can be set to an
// environment variable later on with env.Load().
func Write(key, value, fileName string) error {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileEnvs := parse(fileData)
	fileEnvs[key] = value
	backupFileName := ".back_" + fileName

	_, err = copy(fileName, backupFileName)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	keys := []string{}
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

func copy(source, destination string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return 0, err
	}
	defer destinationFile.Close()

	return io.Copy(sourceFile, destinationFile)
}

func parse(fileData []byte) map[string]string {
	lines := strings.Split(string(fileData), "\n")
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
