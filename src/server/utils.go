package api

import (
	"time"

	"github.com/satori/go.uuid"
)

// NewId: new id, uuid
func NewId() string {
	return uuid.NewV4().String()
}

// NowTimestamp: time stamp
func NowTimestamp() int {
	return int(time.Now().Unix())
}
