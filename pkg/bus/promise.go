package bus

type Promise struct {
	result chan []byte
	err    error
}

func (p *Promise) SetError(err error)             { p.err = err }
func (p *Promise) GetError() error                { return p.err }
func (p *Promise) ErrorToString() string          { return p.err.Error() }
func (p *Promise) IsErrorNil() bool               { return p.err == nil }
func (p *Promise) SetChannel(channel chan []byte) { p.result = channel }
func (p *Promise) GetChannel() chan []byte        { return p.result }
func (p *Promise) Resolve(result []byte)          { p.result <- result }
