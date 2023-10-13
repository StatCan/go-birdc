package gobirdc

import (
	"errors"
	"strings"
)

func (b *BirdClient) Configure(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{CONFIGURE}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}
