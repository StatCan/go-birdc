package gobirdc

import (
	"errors"
	"strings"
)

func (b *BirdClient) DisableProtocol(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{DISABLE_PROTOCOL}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

func (b *BirdClient) EnableProtocol(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{ENABLE_PROTOCOL}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

func (b *BirdClient) RestartProtocol(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{RESTART_PROTOCOL}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

func (b *BirdClient) ReloadProtocol(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{RELOAD_PROTOCOL}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

func (b *BirdClient) Shutdown() (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(DOWN)
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

func (b *BirdClient) GracefulRestart() (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(GRACEFUL_RESTART)
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}
