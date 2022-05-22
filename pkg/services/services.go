package services

type Cached interface {
	LocalClear() error
	Sync() error
	// PartialSync sync data after lastEditTime
	PartialSync(lastEditTime *int32) error
}
