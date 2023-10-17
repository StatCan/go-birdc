package gobirdc

import (
	"errors"
	"strings"
)

// Configure sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc configure [...]`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
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
