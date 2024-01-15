package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("unset env", func(t *testing.T) {
		os.Setenv("TEST", "test")
		env := make(Environment)
		env["TEST"] = EnvValue{
			"test",
			false,
		}
		RunCmd([]string{"ls"}, env)
		cat, ok := os.LookupEnv("TEST")

		require.False(t, ok)
		require.Equal(t, "", cat)
	})

	t.Run("empty env", func(t *testing.T) {
		r := RunCmd([]string{"ls"}, Environment{})
		require.Equal(t, 0, r)
	})
}
