package services

// Synchronizable represents stores data in local, and impl the sync methods.
type Synchronizable interface {
	Clean() error
	Sync() error
	// PartialSync sync data after lastEditTime
	PartialSync(lastEditTime *int32) error
}
