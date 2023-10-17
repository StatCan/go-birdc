package gobirdc

import (
	"errors"
	"strings"
)

// ShowStatus returns the status of the bird daemon
func (b *BirdClient) ShowStatus() (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(SHOW_STATUS)
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowMemory returns the memory usage of the bird daemon
func (b *BirdClient) ShowMemory() (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(SHOW_MEMORY)
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowProtocols returns routing protocols configured on the BIRD daemon
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowProtocols(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_PROTOCOLS}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowInterfaces returns a list or summary of network interfaces on the BIRD host.
// The request issued to the BIRD daemon will be equivalent to `birdc show interfaces`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowInterfaces(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_INTERFACES}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowRoute returns a list or summary of routes on the BIRD host.
// The request issued to the BIRD daemon will be equivalent to `birdc show route`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowRoute(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_ROUTE}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowSymbols returns all known symbolic names in BIRD (ie: protocol, routing tables, etc...)
// The request issued to the BIRD daemon will be equivalent to `birdc show symbols`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowSymbols(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_SYMBOLS}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowBFDSessions returns information about BFD sessions in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show bfd sessions`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowBFDSessions(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_BFD_SESSIONS}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowBabelInterfaces returns information about BABEL interfaces in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show babel interfaces`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowBabelInterfaces(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_BABEL_INTERFACES}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowBabelNeighbors returns information about BABEL neighbors in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show babel neighbors`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowBabelNeighbors(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_BABEL_NEIGHBORS}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowBabelEntries returns information about BABEL prefix entries in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show babel entries`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowBabelEntries(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_BABEL_ENTRIES}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowBabelRoutes returns information about BABEL route entries in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show babel routes`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowBabelRoutes(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_BABEL_ROUTES}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowOSPF returns information about the OSPF protocol configuration in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show ospf`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowOSPF(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_OSPF}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowRIPInterfaces returns information about the RIP interfaces in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show rip interfaces`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowRIPInterfaces(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_RIP_INTERFACES}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowRIPNeighbors returns information about RIP neighbors in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show rip neighbors`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowRIPNeighbors(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_RIP_NEIGHBORS}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}

// ShowStatic returns information about static protocol in BIRD
// The request issued to the BIRD daemon will be equivalent to `birdc show static`
//
// An optional number of arguments can be passed to the configure command that will be forwarded with the bytes sent
// to the BIRD daemon.
//
// Arguments will be sent as-is to the BIRD daemon, so care should be taken to ensure that they are properly escaped.
// Additionally, Arguments should line up with the commandline syntax associated with the official BIRD client (birdc)
func (b *BirdClient) ShowStatic(args ...string) (resp, replyCode []byte, err error) {
	err = b.s.Connect()
	if err != nil {
		return
	}
	defer b.s.Close()

	resp, replyCode, err = b.s.Send(strings.Join(append([]string{SHOW_STATIC}, args...), " "))
	if err != nil {
		return
	}

	if replyCode[0] != '0' {
		err = errors.New("got a non-zero reply-code from the server")
	}

	return
}
