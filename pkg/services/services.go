package services

type Cached interface {
	Invalidate() error
}
