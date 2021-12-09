package services

// Cached ...
type Cached interface {
	Invalidate() error
}
