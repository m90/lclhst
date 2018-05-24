package lclhst

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Wait repeatedly pings localhost at the given port until it's ready.
// If creating the request fails, it will return an error. In case of success,
// nil is returned. In case localhost will never respond, Wait blocks infinitely.
func Wait(port int) error {
	return wait(context.Background(), port)
}

// WaitDuration repeatedly pings localhost at the given port until it's ready.
// If creating the request fails or given duration has passed, it will
// return an error. In case of success, nil is returned.
func WaitDuration(d time.Duration, port int) error {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	return wait(ctx, port)
}

// WaitDeadline repeatedly pings localhost at the given port until it's ready.
// If creating the request fails or given deadline has passed, it will
// return an error. In case of success, nil is returned.
func WaitDeadline(t time.Time, port int) error {
	ctx, cancel := context.WithDeadline(context.Background(), t)
	defer cancel()
	return wait(ctx, port)
}

// WaitContext repeatedly pings localhost at the given port until it's ready.
// If creating the request fails or the context's deadline is exceeded, it will
// return an error. In case of success, nil is returned.
func WaitContext(ctx context.Context, port int) error {
	return wait(ctx, port)
}

func wait(ctx context.Context, port int) error {
	req, reqErr := http.NewRequest(http.MethodHead, fmt.Sprintf("http://localhost:%d/", port), nil)
	if reqErr != nil {
		return reqErr
	}

	schedule := time.NewTimer(0)
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("exceeded deadline waiting for port %v to respond", port)
		case <-schedule.C:
			schedule.Reset(time.Second / 4)
			res, resErr := http.DefaultClient.Do(req)
			if resErr == nil && res.StatusCode == http.StatusOK {
				return nil
			}
		}
	}
}
