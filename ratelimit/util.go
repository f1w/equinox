package ratelimit

import (
	"context"
	"strconv"
	"strings"
	"time"
)

// Checks if the limit usage factor and interval overhead within a valid range.
func ValidateRateLimitOptions(limitUsageFactor float64, intervalOverhead time.Duration) (float64, time.Duration) {
	if limitUsageFactor <= 0.0 || limitUsageFactor > 1.0 {
		limitUsageFactor = 0.99
	}

	if intervalOverhead < 0 {
		intervalOverhead = time.Second
	}

	return limitUsageFactor, intervalOverhead
}

// Waits for the given duration after checking if the context deadline will be exceeded.
func WaitN(ctx context.Context, estimated time.Time, duration time.Duration) error {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(estimated) {
		return ErrContextDeadlineExceeded
	}

	select {
	case <-time.After(duration):
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// Returns the time.Duration in seconds to wait from the Retry-After header, DEFAULT_RETRY_AFTER if not found.
func GetRetryAfterHeader(retryAfterHeader string) time.Duration {
	if retryAfterHeader == "" {
		return DEFAULT_RETRY_AFTER
	}

	delayF, err := strconv.Atoi(retryAfterHeader)
	if err != nil {
		return DEFAULT_RETRY_AFTER
	}

	return time.Duration(delayF) * time.Second
}

// Returns the limit and interval in seconds from a pair of numbers separated by a colon.
func GetNumbersFromPair(pair string) (int, time.Duration) {
	numbers := strings.Split(pair, ":")
	interval, _ := strconv.Atoi(numbers[1])
	limitOrCount, _ := strconv.Atoi(numbers[0])
	return limitOrCount, time.Duration(interval) * time.Second
}
