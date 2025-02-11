package tests

import (
	"context"
	"io"
	"strings"
	"sync"

	"github.com/navidrome/navidrome/utils"
)

func NewMockFFmpeg(data string) *MockFFmpeg {
	return &MockFFmpeg{Reader: strings.NewReader(data)}
}

type MockFFmpeg struct {
	io.Reader
	lock   sync.Mutex
	closed utils.AtomicBool
	Error  error
}

func (ff *MockFFmpeg) Transcode(_ context.Context, _, _ string, _ int) (f io.ReadCloser, err error) {
	if ff.Error != nil {
		return nil, ff.Error
	}
	return ff, nil
}

func (ff *MockFFmpeg) ExtractImage(context.Context, string) (io.ReadCloser, error) {
	if ff.Error != nil {
		return nil, ff.Error
	}
	return ff, nil
}

func (ff *MockFFmpeg) Probe(context.Context, []string) (string, error) {
	if ff.Error != nil {
		return "", ff.Error
	}
	return "", nil
}
func (ff *MockFFmpeg) CmdPath() (string, error) {
	if ff.Error != nil {
		return "", ff.Error
	}
	return "ffmpeg", nil
}

func (ff *MockFFmpeg) Read(p []byte) (n int, err error) {
	ff.lock.Lock()
	defer ff.lock.Unlock()
	return ff.Reader.Read(p)
}

func (ff *MockFFmpeg) Close() error {
	ff.closed.Set(true)
	return nil
}

func (ff *MockFFmpeg) IsClosed() bool {
	return ff.closed.Get()
}
