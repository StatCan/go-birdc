package socket

import (
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

type BirdSocket struct {
	path       string
	bufferSize int
	conn       net.Conn
}

func NewBirdSocket(path string, bufferSize int) *BirdSocket {
	return &BirdSocket{path: path, bufferSize: bufferSize}
}

func (s *BirdSocket) Connect() (err error) {
	s.conn, err = net.Dial("unix", s.path)
	if err != nil {
		return
	}

	// this is to read the connect welcome message into the buffer and flush
	// before any further messages are sent
	buf := make([]byte, s.bufferSize)
	_, err = s.conn.Read(buf[:])

	return
}

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
	err = s.conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		return
	}
	_, err = s.conn.Write([]byte(strings.Trim(string(data), "\n") + "\n"))
	return
}

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

func containsreplyCode(b []byte) bool {
	return replyCodeExpr.Match(b)
}
