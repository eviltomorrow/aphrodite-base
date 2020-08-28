package tools

import (
	"fmt"
	"os"
	"syscall"
)

// FileLock file lock
type FileLock interface {
	Release() error
}

// UnixFileLock file lock
type UnixFileLock struct {
	f *os.File
}

// Release release
func (fl *UnixFileLock) Release() error {
	if fl == nil {
		return nil
	}
	if err := setFileLock(fl.f, false, false); err != nil {
		return err
	}
	return fl.f.Close()
}

// NewFileLock new file lock
func NewFileLock(path string, readOnly bool) (fl FileLock, err error) {
	var flag int
	if readOnly {
		flag = os.O_RDONLY
	} else {
		flag = os.O_RDWR | os.O_TRUNC
	}
	f, err := os.OpenFile(path, flag, 0)
	if os.IsNotExist(err) {
		f, err = os.OpenFile(path, flag|os.O_CREATE|os.O_TRUNC, 0644)
	}
	if err != nil {
		return
	}

	var pid = os.Getpid()
	f.WriteString(fmt.Sprintf("%d", pid))

	err = setFileLock(f, readOnly, true)
	if err != nil {
		f.Close()
		return
	}
	fl = &UnixFileLock{f: f}
	return
}

func setFileLock(f *os.File, readOnly, lock bool) error {
	how := syscall.LOCK_UN
	if lock {
		if readOnly {
			how = syscall.LOCK_SH
		} else {
			how = syscall.LOCK_EX
		}
	}
	return syscall.Flock(int(f.Fd()), how|syscall.LOCK_NB)
}
