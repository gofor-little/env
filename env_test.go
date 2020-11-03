package env_test

import (
	"os"
	"testing"

	"github.com/gofor-little/env"
)

func TestLoad(t *testing.T) {
	if err := env.Load("test-data/test-1.env", "test-data/test-2.env"); err != nil {
		t.Fatalf("failed to load .env files: %v", err)
	}

	if env.Get("PUBLIC_KEY", "") != os.Getenv("PUBLIC_KEY") {
		t.Fatalf("unexpected result from PUBLIC_KEY environment variable: %s", os.Getenv("PUBLIC_KEY"))
	}

	if env.Get("PRIVATE_KEY", "") != os.Getenv("PRIVATE_KEY") {
		t.Fatalf("unexpected result from PRIVATE_KEY environment variable: %s", os.Getenv("PRIVATE_KEY"))
	}

	if env.Get("DB_NAME", "") != os.Getenv("DB_NAME") {
		t.Fatalf("unexpected result from DB_NAME environment variable: %s", os.Getenv("DB_NAME"))
	}

	if env.Get("DB_PASSWORD", "") != os.Getenv("DB_PASSWORD") {
		t.Fatalf("unexpected result from DB_PASSWORD environment variable: %s", os.Getenv("DB_PASSWORD"))
	}
}

func TestWrite(t *testing.T) {
	if err := env.Write("PUBLIC_KEY_NEW", "\"public\\nkey\\nnew\"", "test-data/test-1.env", true); err != nil {
		t.Fatalf("failed to write to .env file: %v", err)
	}

	if err := env.Write("DB_PASSWORD", "db_password_new", "test-data/test-2.env", true); err != nil {
		t.Fatalf("failed to write to .env file: %v", err)
	}
}

func TestGet(t *testing.T) {
	if err := env.Load("test-data/test-1.env", "test-data/test-2.env"); err != nil {
		t.Fatalf("failed to load .env files: %v", err)
	}

	value := env.Get("DB_NAME", "db_name_default")
	if value != "db_name" {
		t.Fatalf("unexpected result from DB_NAME environment variable: %s", value)
	}

	value = env.Get("FAKE_KEY", "db_name_default")
	if value != "db_name_default" {
		t.Fatalf("unexpected result from FAKE_KEY environment variable: %s", value)
	}
}

func TestMustGet(t *testing.T) {
	if err := env.Load("test-data/test-1.env", "test-data/test-2.env"); err != nil {
		t.Fatalf("failed to load .env files: %v", err)
	}

	value, err := env.MustGet("DB_NAME")
	if err != nil {
		t.Fatalf("failed to get environment variable for key DB_NAME: %v", err)
	}

	if value != "db_name" {
		t.Fatalf("unexpected result from DB_NAME environment variable: %s", value)
	}

	_, err = env.MustGet("FAKE_KEY")
	if err == nil {
		t.Fatalf("failed to get environment variable for key FAKE_KEY: %v", err)
	}
}

func TestSet(t *testing.T) {
	if err := env.Load("test-data/test-1.env", "test-data/test-2.env"); err != nil {
		t.Fatalf("failed to load .env files: %v", err)
	}

	if err := env.Set("DB_NAME", "db_name_override"); err != nil {
		t.Fatalf("failed to set environment variable for key DB_NAME: %v", err)
	}

	value := os.Getenv("DB_NAME")
	if value != "db_name_override" {
		t.Fatalf("unexpected result from DB_NAME environment variable: %s", value)
	}
}
