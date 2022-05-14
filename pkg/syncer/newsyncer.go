package syncer

import "github.com/alswl/go-toodledo/pkg/models"

type Progress interface {
	Start(int)
	Update(int)
	Finish()
}

type SyncItemType string
type SyncStatus string

var (
	SyncItemTypeCreate SyncItemType = "CREATE"
	SyncItemTypePatch  SyncItemType = "PATCH"
	SyncItemTypeDelete SyncItemType = "DELETE"

	SyncStatusSynced  SyncStatus = "SYNCED"
	SyncStatusSyncing SyncStatus = "SYNCING"
	SyncStatusError   SyncStatus = "ERROR"
)

type SyncItem struct {
	Item *models.Task
	Type SyncItemType
}

type Syncer2 interface {
	Run(stopCh chan struct{}) error
	Sync(diffs []*SyncItem, progress Progress) (int, int, []*models.Task, error)
	Status() (SyncStatus, error)
}
