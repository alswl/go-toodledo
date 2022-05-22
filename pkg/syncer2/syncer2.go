package syncer2

import "github.com/alswl/go-toodledo/pkg/models"

const (
	SyncItemTypeCreate SyncItemType = "CREATE"
	SyncItemTypeUpdate SyncItemType = "UPDATE"
	SyncItemTypeDelete SyncItemType = "DELETE"

	SyncStatusSynced  SyncStatus = "SYNCED"
	SyncStatusSyncing SyncStatus = "SYNCING"
	SyncStatusError   SyncStatus = "ERROR"
)

type SyncEvent struct {
	Item *models.Task
	Type SyncItemType
}

// Syncer2 is two-way sync
type Syncer2[T any] interface {
	Run(stopCh chan struct{}) error
	//Sync(diffs []*SyncEvent, callback func(progress Progress) error) (int, int, []*models.Task, error)

	PostEvent(event SyncEvent) error
	Status() string

	SubscribeStatus(func() (string, error)) error
	AddFun(item T) error
	DeleteFun(item T) error
	UpdateFun(item T) error
}

type TaskSyncer Syncer2[models.Task]
