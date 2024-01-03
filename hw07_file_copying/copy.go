package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer file.Close()

	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset >= fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileCopy, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer fileCopy.Close()

	size := fileInfo.Size() - offset
	if limit > 0 && size > limit {
		size = limit
	}

	bar := pb.Full.Start64(size)
	barReader := bar.NewProxyReader(file)

	if limit == 0 && offset == 0 {
		if _, err := io.Copy(fileCopy, barReader); err != nil {
			return err
		}
		return nil
	}

	result := make([]byte, size)

	if _, err = file.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	if _, err = barReader.Read(result); err != nil {
		return err
	}

	if _, err = fileCopy.Write(result); err != nil {
		return err
	}

	bar.Finish()

	return nil
}
