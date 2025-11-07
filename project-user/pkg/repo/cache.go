package repo

import (
	"time"
)

type Cache interface {
	Put(key, value string, _ time.Duration) error
	Get(key string) (string, error)
}
