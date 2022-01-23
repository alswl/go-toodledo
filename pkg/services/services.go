package services

// Cached ...
type Cached interface {
	LocalClear() error
	Sync() error
}
