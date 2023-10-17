package socket

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

var replyCodeExpr *regexp.Regexp

func init() {
	// https://gitlab.nic.cz/labs/bird/-/blob/master/doc/reply_codes
	replyCodeExpr = regexp.MustCompile(`(?m)^([089][0-9]{3})`)
}

// BirdSocket represents a socket connection to bird daemon
//
// `path` is the unix socket path
//
// `bufferSize` is the size of the read buffer
//
// `conn` is the actual socket connection
type BirdSocket struct {
	path       string
	bufferSize int
	conn       net.Conn
}

// NewBirdSocket creates a new BirdSocket
//
// `path` is the unix socket path
//
// `bufferSize` is the size of the read buffer
func NewBirdSocket(path string, bufferSize int) *BirdSocket {
	return &BirdSocket{path: path, bufferSize: bufferSize}
}

// Connect opens the unix socket connection
func (s *BirdSocket) Connect() (err error) {
	s.conn, err = net.Dial("unix", s.path)
	if err != nil {
		return
	}

	// throw away the welcome message
	buf := make([]byte, s.bufferSize)
	_, err = s.conn.Read(buf[:])

	return
}

// Close closes the unix socket connection
//
// **NOTE** it is important to close the socket connection since,
// in the event that there is still data waiting to be transmitted over the connection,
// close will try to complete the transmission properly
func (s *BirdSocket) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *BirdSocket) Send(data string) (resp []byte, replyCode []byte, err error) {
	err = s.write([]byte(data))
	if err != nil {
		return
	}
	return s.read()
}

func (s *BirdSocket) write(data []byte) (err error) {
	if s.conn == nil {
		err = fmt.Errorf("unable to write to socket, connection is not open. Ensure that s.Connect() is called first. %s", s.conn)
		return
	}

	err = s.conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		return
	}
	_, err = s.conn.Write([]byte(strings.Trim(string(data), "\n") + "\n"))
	return
}

// read reads data from the socket connection and captures the replyCode that indicated the end of transmissions
// and any associated response or error
//
// For more information on reply codes, see https://gitlab.nic.cz/labs/bird/-/blob/master/doc/reply_codes
//
// `resp` is the raw response from BIRD
//
// `replyCode` is the reply code associated with the response
func (s *BirdSocket) read() (resp []byte, replyCode []byte, err error) {
	err = s.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		return
	}

	resp = make([]byte, 0)
	data := make([]byte, s.bufferSize)

	done := false
	for !done {
		var n int
		n, err = s.conn.Read(data)
		if err != nil {
			return
		}

		resp = append(resp, data[:n]...)
		done = containsreplyCode(data[:n])
	}

	replyCode = replyCodeExpr.Find(resp)
	return
}

// containsreplyCode checks if the given byte slice contains a valid reply code
// indicitive of the end of transmission
func containsreplyCode(b []byte) bool {
	return replyCodeExpr.Match(b)
}
