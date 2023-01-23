package service

import "time"

type DaysCounter interface {
	CalculateDaysGoneOrLeft(time time.Time) string
}

type Service struct {
	DaysCounter
}

func NewService() *Service {
	return &Service{DaysCounter: NewDCounter()}
}
