//go:build integration

package fetcher

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestFetcherProgress(t *testing.T) {
	log := logrus.New()
	// fn present the fetch function
	fn := func(sd fetchers.StatusDescriber) error {
		ticker := time.NewTicker(time.Second * 1)
		timeout := time.After(time.Second * 10)
		i := 0
		for {
			select {
			case <-ticker.C:
				i += 1
				sd.SetProgress(i, 10)
				t.Log("tick")
			case <-timeout:
				t.Log("now timeout.")
				return nil
			}

		}
		return nil
	}
	sd := fetchers.NewNoOpStatusDescriber()
	f := fetchers.NewSimpleFetcher(log, 5*time.Minute, fn, nil)
	ctx := context.Background()
	f.Start(ctx)

	// print progress interval
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				current, total := sd.Progress()
				t.Logf("current: %d, total: %d", current, total)
			}
		}
	}()

	// stop all
	time.Sleep(time.Second * 30)
	f.Stop()
}
