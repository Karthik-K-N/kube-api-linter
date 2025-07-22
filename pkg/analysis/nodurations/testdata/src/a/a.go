package a

import "time"

type Durations struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt64 int64

	InvalidDurations time.Duration
}
