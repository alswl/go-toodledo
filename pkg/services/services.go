package services

// LocalStorage represents stores data in local, and impl the sync methods.
type LocalStorage interface {
	LocalClear() error
	Sync() error
	// PartialSync sync data after lastEditTime
	PartialSync(lastEditTime *int32) error
}
