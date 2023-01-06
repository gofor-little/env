## A package for managing .env files and environment variables

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/env?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/env)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/env/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/gofor-little/env/ci.yaml?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/env)](https://goreportcard.com/report/github.com/gofor-little/env)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/env)](https://pkg.go.dev/github.com/gofor-little/env)

### Introduction
* Read from and write to .env files
* Get and set environment variables
* No dependencies outside the standard library

### Example
```go
package main

import "github.com/gofor-little/env"

func main() {
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load("FILE_PATH"); err != nil {
		panic(err)
	}

	// Write a key-value pair to an .env file and call env.Set on it.
	if err := env.Write("KEY", "VALUE", "FILE_PATH", true); err != nil {
		panic(err)
	}

	// Get an environment variable's value with a default backup value.
	value := env.Get("KEY", "DEFAULT_VALUE")

	// Get an environment variable's value, receiving an error if it is not set or is empty.
	value, err := env.MustGet("KEY")
	if err != nil {
		panic(err)
	}

	// Set an environment variable locally.
	if err := env.Set("KEY", "VALUE"); err != nil {
		panic(err)
	}
}
```

### Testing
Run ```go test ./...``` in the root directory.
