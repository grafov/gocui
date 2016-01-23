package gocui

import "errors"

var (
	// ErrQuit is used to decide if the MainLoop finished succesfully.
	ErrQuit = errors.New("quit")

	// ErrUnknownView allows to assert if a View must be initialized.
	ErrUnknownView = errors.New("unknown view")

	// ErrViewSizeTooSmall is appeared when resized view is unnormally small.
	ErrViewSizeTooSmall = errors.New("view size can't be such small")
)
