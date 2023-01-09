// Package editor provides a simple interface for launching a text editor.
// it was inspired by kubectl edit and comes from
// https://github.com/yuuu/go-editor/blob/master/editor.go with some patch
package editor

import (
	"fmt"
	"os"
	"os/exec"
)

const DefaultEditor = "vi"

type Editor struct {
	cmd string
}

type Error struct {
	msg string
	err error
}

func NewDefaultEditor() (*Editor, error) {
	return NewEditor(DefaultEditor)
}

func NewEditor(cmd string) (*Editor, error) {
	var err error

	_, err = exec.LookPath(cmd)
	if err != nil {
		envEditor := os.Getenv("EDITOR")
		if envEditor != "" {
			cmd = envEditor
		} else {
			cmd = DefaultEditor
		}
		_, err = exec.LookPath(cmd)
		if err != nil {
			return nil, err
		}
	}

	editor := Editor{cmd}
	return &editor, nil
}

func (editor *Editor) Launch(filepath string) error {
	var cmdErr error

	// nolint:gosec
	cmd := exec.Command(editor.cmd, filepath)
	if cmd == nil {
		err := Error{"Invalid command", nil}
		return &err
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr = cmd.Start()
	if cmdErr != nil {
		err := Error{"Failed to Start", cmdErr}
		return &err
	}

	cmdErr = cmd.Wait()
	if cmdErr != nil {
		err := Error{"Failed to Wait", cmdErr}
		return &err
	}

	return nil
}

func (err *Error) Error() string {
	return fmt.Errorf("%s: %w", err.msg, err.err).Error()
}
