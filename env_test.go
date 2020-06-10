package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	env "github.com/taliesinmillhouse/little-env"
)

func TestLoad(t *testing.T) {
	if !assert.NoError(t, env.Load("test-1.env", "test-2.env"), "failed to load .env files") {
		return
	}

	if !assert.Equal(t, "public_key", os.Getenv("PUBLIC_KEY"), "unexpected result from PUBLIC_KEY environment variable") {
		return
	}

	if !assert.Equal(t, "private_key", os.Getenv("PRIVATE_KEY"), "unexpected result from PUBLIC_KEY environment variable") {
		return
	}

	if !assert.Equal(t, "db_name", os.Getenv("DB_NAME"), "unexpected result from PUBLIC_KEY environment variable") {
		return
	}

	if !assert.Equal(t, "db_password", os.Getenv("DB_PASSWORD"), "unexpected result from PUBLIC_KEY environment variable") {
		return
	}
}

func TestWrite(t *testing.T) {
	if !assert.NoError(t, env.Write("PUBLIC_KEY_NEW", "public_key_new", "test-1.env"), "failed to write to .env file") {
		return
	}
}
