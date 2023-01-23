package service

import (
	"fmt"
	"time"
)

type DCounter struct {
}

func NewDCounter() *DCounter {
	return &DCounter{}
}

func (dc *DCounter) CalculateDaysGoneOrLeft(t time.Time) string {
	var dur int64
	year := t.Year()
	if year > 2023 {
		dur = int64(time.Until(t).Hours())
		return fmt.Sprintf("Days left: %d", dur/24)
	} else {
		dur = int64(time.Since(t).Hours())
		return fmt.Sprintf("Days gone: %d", dur/24)
	}
}
