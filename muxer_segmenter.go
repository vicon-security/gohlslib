package gohlslib

import (
	"time"
)

type muxerSegmenter interface {
	close()
	reset()
	writeAV1(time.Time, time.Duration, [][]byte, bool, bool) error
	writeVP9(time.Time, time.Duration, []byte, bool, bool) error
	writeH26x(time.Time, time.Duration, [][]byte, bool, bool, bool) error
	writeOpus(time.Time, time.Duration, [][]byte) error
	writeMPEG4Audio(time.Time, time.Duration, [][]byte) error
	queueToStop()
	isStopped() bool
	startSegments()
}
