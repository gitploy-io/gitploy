package interactor

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (i *Interactor) runAutoUnlock(stop <-chan struct{}) {
	ctx := context.Background()

	ticker := time.NewTicker(time.Minute)
L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ls, err := i.ListExpiredLocksLessThanTime(ctx, t)
			if err != nil {
				i.log.Error("It has failed to read expired locks.", zap.Error(err))
				continue
			}

			for _, l := range ls {
				i.DeleteLock(ctx, l)
				i.log.Debug("Delete the expired lock.", zap.Int("id", l.ID))
			}
		}
	}
}
