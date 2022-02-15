package services

// Cached ...
type Cached interface {
	LocalClear() error
	Sync() error
	PartialSync(lastEditTime *int32) error
}
