package challenge

import (
	"context"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/config"
	"github.com/desertbit/timer"
)

// SetupWorkers start challenge workers
func SetupWorkers(ctx context.Context) {
	go startAcceptNew(ctx)
	go startProcessAccepted(ctx)
	go startCommitProcessed(ctx)
}

func startCommitProcessed(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After((time.Duration(config.Configuration.ChallengeResolveFreq) * time.Second):
			commitProcessed(ctx)
		}
	}
}

func startProcessAccepted(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.After(time.Duration(config.Configuration.ChallengeResolveFreq) * time.Second):
			processAccepted(ctx)
		}
	}
}

// startAcceptNew
func startAcceptNew(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(config.Configuration.ChallengeResolveFreq) * time.Second):
			acceptChallenges(ctx)
		}
	}
}
