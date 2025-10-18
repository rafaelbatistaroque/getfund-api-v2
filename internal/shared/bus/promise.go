package shared_bus

type Promise struct {
	result chan []byte
	err    error
}

// SetError sets the error of the promise.
func (p *Promise) SetError(err error)             { p.err = err }
// GetError returns the error of the promise.
func (p *Promise) GetError() error                { return p.err }
// ErrorToString returns the error of the promise as a string.
func (p *Promise) ErrorToString() string          { return p.err.Error() }
// HasError returns true if the promise has an error.
func (p *Promise) HasError() bool                 { return p.err != nil }
// SetChannel sets the channel for the promise result.
func (p *Promise) SetChannel(channel chan []byte) { p.result = channel }
// GetChannel returns the channel for the promise result.
func (p *Promise) GetChannel() chan []byte        { return p.result }
// Resolve resolves the promise with a result.
func (p *Promise) Resolve(result []byte)          { p.result <- result }
