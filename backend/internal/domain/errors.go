package domain

import "errors"

var (
	ErrQuoteNotFound  = errors.New("quote not found")
	ErrUpstream       = errors.New("upstream source unavailable")
	ErrNotConfigured  = errors.New("provider not configured")
	ErrNotImplemented = errors.New("provider not implemented yet")
)
