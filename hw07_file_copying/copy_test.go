package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	const parentFile = "./testdata/input.txt"
	const parentFileInvalid = "./testdata/invalid_offset.txt"

	err := os.Mkdir("tmp", os.ModePerm)
	require.NoError(t, err)
	t.Run("all file", func(t *testing.T) {
		err := Copy(parentFile, "tmp/test.txt", 0, 0)
		require.NoError(t, err)

		fileCopy, err := os.ReadFile("tmp/test.txt")
		require.NoError(t, err)

		fileInput, err := os.ReadFile(parentFile)
		require.NoError(t, err)

		require.Equal(t, string(fileInput), string(fileCopy))
	})
	t.Run("Invalid offset", func(t *testing.T) {
		err := Copy(parentFileInvalid, "tmp/invalid.txt", 30, 0)
		require.Error(t, err, ErrOffsetExceedsFileSize)
	})
	t.Run("Unsupported file", func(t *testing.T) {
		err := Copy("./unsupfile.txt", "tmp/invalid.txt", 30, 0)
		require.Error(t, err, ErrUnsupportedFile)
	})
	t.Run("out offset=0 limit=10", func(t *testing.T) {
		err := Copy(parentFile, "tmp/test1.txt", 0, 10)
		require.NoError(t, err)

		fileCopy, err := os.ReadFile("tmp/test1.txt")
		require.NoError(t, err)

		fileInput, err := os.ReadFile("./testdata/out_offset0_limit10.txt")
		require.NoError(t, err)

		require.Equal(t, string(fileInput), string(fileCopy))
	})
	t.Run("out offset=100 limit=1000", func(t *testing.T) {
		err := Copy(parentFile, "tmp/test2.txt", 100, 1000)
		require.NoError(t, err)

		fileCopy, err := os.ReadFile("tmp/test2.txt")
		require.NoError(t, err)

		fileInput, err := os.ReadFile("./testdata/out_offset100_limit1000.txt")
		require.NoError(t, err)

		require.Equal(t, string(fileInput), string(fileCopy))
	})
	t.Run("out offset=6000 limit=1000", func(t *testing.T) {
		err := Copy(parentFile, "tmp/test9.txt", 6000, 1000)
		require.NoError(t, err)

		fileCopy, err := os.ReadFile("tmp/test9.txt")
		require.NoError(t, err)

		fileInput, err := os.ReadFile("./testdata/out_offset6000_limit1000.txt")
		require.NoError(t, err)

		require.Equal(t, string(fileInput), string(fileCopy))
	})
	t.Cleanup(func() {
		err := os.RemoveAll("tmp/")
		require.NoError(t, err)
	})
}
