//go:build integration || tags
// +build integration tags

package informers

import (
	"context"
	"testing"
	"time"

	"github.com/alswl/go-toodledo/pkg/common/informers"
	"github.com/sirupsen/logrus"
)

func TestReflector(t *testing.T) {
	log := logrus.StandardLogger()
	r := informers.NewReflector(log, 10*time.Second, func(ctx context.Context, lastSynced informers.U) ([]informers.T, informers.U, error) {
		var last = int64(42)
		return []informers.T{""}, &last, nil
	})
	stopCh := make(chan struct{})
	// forever loop
	r.Run(stopCh)
}

func TestReflectorStop(t *testing.T) {
	log := logrus.StandardLogger()
	r := informers.NewReflector(log, 10*time.Second, func(ctx context.Context, lastSynced informers.U) ([]informers.T, informers.U, error) {
		select {
		case <-ctx.Done():
		case <-time.After(2 * time.Second):
		}
		var last = int64(42)

		return []informers.T{""}, &last, nil
	})
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(15 * time.Second)
		stopCh <- struct{}{}
	}()
	r.Run(stopCh)
}
