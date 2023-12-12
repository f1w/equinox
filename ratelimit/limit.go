package ratelimit

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Limit represents a collection of buckets and the type of limit (application or method).
type Limit struct {
	limitType  string
	buckets    []*Bucket
	retryAfter time.Duration
	mutex      sync.Mutex
}

func NewLimit(limitType string) *Limit {
	return &Limit{
		buckets:    make([]*Bucket, 0),
		limitType:  limitType,
		retryAfter: 0,
		mutex:      sync.Mutex{},
	}
}

// Checks if any of the buckets provided are rate limited, and if so, blocks until the next reset.
func (l *Limit) checkBuckets(ctx context.Context, logger zerolog.Logger, route any, methodID string) error {
	if l.retryAfter > 0 {
		err := WaitN(ctx, time.Now().Add(l.retryAfter), l.retryAfter)
		if err != nil {
			return err
		}
		l.retryAfter = 0
	}

	var limited []*Bucket
	for _, bucket := range l.buckets {
		if bucket.isRateLimited() {
			limited = append(limited, bucket)
		}
	}

	for i := len(limited) - 1; i >= 0; i-- {
		bucket := limited[i]
		logger.Warn().
			Any("route", route).
			Str("method_id", methodID).
			Str("limit_type", l.limitType).
			Object("bucket", bucket).
			Msg("Rate limited")
		bucket.mutex.Lock()
		defer bucket.mutex.Unlock()
		bucket.check()
		err := WaitN(ctx, bucket.next, time.Until(bucket.next))
		if err != nil {
			return err
		}
		bucket.check()
		bucket.tokens--
	}

	return nil
}

// Checks if the limits given in the header match the current buckets.
func (l *Limit) limitsDontMatch(limitHeader string, limitOffset int) bool {
	if limitHeader == "" {
		return false
	}
	limits := strings.Split(limitHeader, ",")
	if len(l.buckets) != len(limits) {
		return true
	}
	for i, pair := range limits {
		if l.buckets[i] == nil {
			return true
		}
		limit, interval := getNumbersFromPair(pair)
		if l.buckets[i].limit != limit-limitOffset || l.buckets[i].interval != time.Duration(interval)*time.Second {
			return true
		}
	}
	return false
}

func (l *Limit) setDelay(delay time.Duration) {
	l.mutex.Lock()
	l.retryAfter = delay
	l.mutex.Unlock()
}
