package env_test

import (
	"os"
	"testing"

	"github.com/gofor-little/env"
	"github.com/matryer/is"
)

func TestLoad(t *testing.T) {
	is := is.New(t)
	is.NoErr(env.Load("test-data/test-1.env", "test-data/test-2.env"))

	is.True(env.Get("PUBLIC_KEY", "") == os.Getenv("PUBLIC_KEY"))
	is.True(env.Get("PRIVATE_KEY", "") == os.Getenv("PRIVATE_KEY"))
	is.True(env.Get("DB_NAME", "") == os.Getenv("DB_NAME"))
	is.True(env.Get("DB_PASSWORD", "") == os.Getenv("DB_PASSWORD"))
}

func TestWrite(t *testing.T) {
	is := is.New(t)

	is.NoErr(env.Write("PUBLIC_KEY_NEW", "\"public\\nkey\\nnew\"", "test-data/test-1.env", true))
	is.NoErr(env.Write("DB_PASSWORD", "db_password_new", "test-data/test-2.env", true))
}

func TestGet(t *testing.T) {
	is := is.New(t)
	is.NoErr(env.Load("test-data/test-1.env", "test-data/test-2.env"))

	is.True(env.Get("DB_NAME", "db_name_default") == "db_name")
	is.True(env.Get("FAKE_KEY", "db_name_default") == "db_name_default")
}

func TestMustGet(t *testing.T) {
	is := is.New(t)
	is.NoErr(env.Load("test-data/test-1.env", "test-data/test-2.env"))

	value, err := env.MustGet("DB_NAME")
	is.NoErr(err)
	is.True(value == "db_name")

	_, err = env.MustGet("FAKE_KEY")
	is.True(err != nil)
}

func TestSet(t *testing.T) {
	is := is.New(t)
	is.NoErr(env.Load("test-data/test-1.env", "test-data/test-2.env"))

	is.NoErr(env.Set("DB_NAME", "db_name_override"))

	value := os.Getenv("DB_NAME")
	is.True(value == "db_name_override")
}
