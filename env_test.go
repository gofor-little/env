package env_test

import (
	"os"
	"testing"

	"github.com/gofor-little/env"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	require.NoError(t, env.Load("test-data/test-1.env", "test-data/test-2.env", "test-data/test-3.env"))

	require.True(t, env.Get("PUBLIC_KEY", "") == os.Getenv("PUBLIC_KEY"))
	require.True(t, env.Get("PRIVATE_KEY", "") == os.Getenv("PRIVATE_KEY"))
	require.True(t, env.Get("DB_NAME", "") == os.Getenv("DB_NAME"))
	require.True(t, env.Get("DB_PASSWORD", "") == os.Getenv("DB_PASSWORD"))
	require.True(t, env.Get("SPECIAL_CHARACTERS", "") == os.Getenv("SPECIAL_CHARACTERS"))
}

func TestWrite(t *testing.T) {
	require.NoError(t, env.Write("PUBLIC_KEY_NEW", "\"public\\nkey\\nnew\"", "test-data/test-1.env", true))
	require.NoError(t, env.Write("DB_PASSWORD", "db_password_new", "test-data/test-2.env", true))
}

func TestGet(t *testing.T) {
	require.NoError(t, env.Load("test-data/test-1.env", "test-data/test-2.env", "test-data/test-3.env"))

	require.True(t, env.Get("DB_NAME", "db_name_default") == "db_name")
	require.True(t, env.Get("SPECIAL_CHARACTERS", "") == "special=characters")
	require.True(t, env.Get("FAKE_KEY", "db_name_default") == "db_name_default")
	require.True(t, env.Get("COMMENTED_ENV", "") == "")
	require.True(t, env.Get("INLINE_COMMENT_ENV", "") == "inline_comment_env")
	require.True(t, env.Get("IGNORE_HASHTAG_ENV", "") == "#ignore_hashtag_env#")
}

func TestMustGet(t *testing.T) {
	require.NoError(t, env.Load("test-data/test-1.env", "test-data/test-2.env"))

	value, err := env.MustGet("DB_NAME")
	require.NoError(t, err)
	require.True(t, value == "db_name")

	value, err = env.MustGet("SPECIAL_CHARACTERS")
	require.NoError(t, err)
	require.True(t, value == "special=characters")

	_, err = env.MustGet("FAKE_KEY")
	require.True(t, err != nil)
}

func TestSet(t *testing.T) {
	require.NoError(t, env.Load("test-data/test-1.env", "test-data/test-2.env"))
	require.NoError(t, env.Set("DB_NAME", "db_name_override"))

	value := os.Getenv("DB_NAME")
	require.True(t, value == "db_name_override")
}
