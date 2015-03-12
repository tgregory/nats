// +build js
package nats

import (
	"bufio"
	"errors"
)

// Read a control line and process the intended op.
func (nc *Conn) readOp(c *control) error {
	if nc.isClosed() {
		return ErrConnectionClosed
	}
	br := bufio.NewReaderSize(nc.conn, defaultBufSize)
	line, err := br.ReadString('\n')
	if err != nil {
		return err
	}
	parseControl(line, c)
	return nil
}
