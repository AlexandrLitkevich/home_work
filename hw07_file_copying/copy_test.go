package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const path = "/Users/a.litkevich/Documents/Learn/Otus/home_work/hw07_file_copying/testdata/input.txt"

func TestCopy(t *testing.T) {
	t.Run("testdata/input.txt", func(t *testing.T) {
		err := Copy(path, "tmp/test", 0, 0)

		require.NoError(t, err)

	})
}
