package service

import (
	"time"

	"github.com/cameronbrill/brill-wtf-go/model"
)

type NewLinkOption func(*model.Link) error

func WithShortURL(want string) NewLinkOption {
	return func(l *model.Link) error {
		l.Want = want
		return nil
	}
}

func WithTTL(ttl time.Duration) NewLinkOption {
	return func(l *model.Link) error {
		l.TTL = ttl
		return nil
	}
}

type ServiceOption func(*s) error
