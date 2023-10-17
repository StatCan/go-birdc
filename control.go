package gobirdc

import (
	"errors"
	"strings"
)

// DisableProtocol sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc disable [...]`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
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

// EnableProtocol sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc enable [...]`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
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

// RestartProtocol sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc restart [...]`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
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

// ReloadProtocol sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc reload [...]`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
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

// Shutdown sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc down`
//
// Shutdown will send a shutdown signal to the BIRD daemon.
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

// GracefulRestart sends the equivalent instruction to the BIRD Unix Domain Socket as `birdc graceful restart`
//
// GracefulRestart will send a signal the BIRD daemon to go down gracefully
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
