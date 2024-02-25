package util

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	tmpDir   = "/tmp"
	liveFile = "/tmp/live"
)

func Create() (*os.File, error) {
	return os.Create(getLiveFile())
}

func Remove() error {
	return os.Remove(getLiveFile())
}

func Exists() bool {

	if _, err := os.Stat(getLiveFile()); errors.Is(err, os.ErrNotExist) {
		_, _ = io.WriteString(os.Stderr, fmt.Sprintf("failed to stat live file. %s", err))
		return false
	} else {
		return true
	}
}

func getLiveFile() string {

	podUid := os.Getenv("POD_UID")
	if len(podUid) == 0 {
		return liveFile
	}

	return fmt.Sprintf("%s/%s", tmpDir, podUid)
}
