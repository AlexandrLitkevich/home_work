package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read dir", func(t *testing.T) {
		envs, err := ReadDir("./testdata/env/")
		require.NoError(t, err)
		t.Log(envs)
		require.False(t, envs["BAR"].NeedRemove)
		require.False(t, envs["HELLO"].NeedRemove)
		require.False(t, envs["FOO"].NeedRemove)

		require.True(t, envs["EMPTY"].NeedRemove)
		require.True(t, envs["UNSET"].NeedRemove)
	})
}
