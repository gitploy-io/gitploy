package stream

import (
	"go.uber.org/zap"
)

type (
	Stream struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewStream(i Interactor) *Stream {
	return &Stream{
		i:   i,
		log: zap.L().Named("stream"),
	}
}
