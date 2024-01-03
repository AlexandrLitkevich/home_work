package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
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

	if offset >= fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileCopy, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer fileCopy.Close()
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(file)

	if limit == 0 {
		_, err := io.Copy(fileCopy, barReader)
		if err != nil {
			return err
		}
		return nil
	}

	result := make([]byte, limit)

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	n, err := barReader.Read(result)
	if err != nil {
		return err
	}

	_, err = fileCopy.Write(result[:n])
	if err != nil {
		return err
	}
	bar.Finish()

	return nil
}
