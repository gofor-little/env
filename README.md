## A package for managing .env files and environment variables
* Read from and write to .env files
* Get and set environment variables

# Example
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
