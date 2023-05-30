package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

var _ Hook = (*hook)(nil)

// Hook is a graceful shutdown way to close http or grpc connection.
// default listen with signals of SIGINT and SIGTERM
type Hook interface {
	// WithSignals: listen with signals
	WithSignals(signals ...syscall.Signal) Hook

	// Close: register shutdown handlers
	Close(funcs ...func())
}

type hook struct {
	ch chan os.Signal
}

// NewHook create a Hook instance
func NewHook() Hook {
	hook := &hook{
		ch: make(chan os.Signal, 1),
	}

	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

// WithSignals: register more signals into hook
func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.ch, s)
	}

	return h
}

// Close: register handlers to hook
// handlers will be executed when receive signals
func (h *hook) Close(funcs ...func()) {
	select {
	case <-h.ch: // if no signal received, will stuck here
	}
	signal.Stop(h.ch)

	for _, f := range funcs {
		f()
	}
}
