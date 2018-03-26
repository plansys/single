// package single provides a mechanism to ensure, that only one instance of a program is running

package single

import (
	"errors"
	"os"
)

var (
	// ErrAlreadyRunning
	ErrAlreadyRunning = errors.New("the program is already running")
	//
	Lockfile string
)

// Single represents the name and the open file descriptor
type Single struct {
	name string
	tempdir string
	file *os.File
}

// New creates a Single instance
func New(name string) *Single {
	return &Single{
		name: name,
		tempdir: "",
	}
}

func NewWithTempDir(name, tempdir string) *Single {
	return &Single{
		name: name,
		tempdir: tempdir,
	}
}
