// +build !js

package nats

import (
	"bufio"
	"errors"
	"unsafe"
)

// Read a control line and process the intended op.
func (nc *Conn) readOp(c *control) error {
	if nc.isClosed() {
		return ErrConnectionClosed
	}
	br := bufio.NewReaderSize(nc.conn, defaultBufSize)
	b, pre, err := br.ReadLine()
	if err != nil {
		return err
	}
	if pre {
		// FIXME: Be more specific here?
		return errors.New("nats: Line too long")
	}
	// Do straight move to string rep.
	line := *(*string)(unsafe.Pointer(&b))
	parseControl(line, c)
	return nil
}
