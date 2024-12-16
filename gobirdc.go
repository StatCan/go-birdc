package gobirdc

import (
	"github.com/StatCan/go-birdc/socket"
)

// A client wrapper for gobirdc which provides access to a number of
// thethe functions/methods available in the birdc binary
type BirdClient struct {
	s *socket.BirdSocket
}

// BirdClientOptions describe configuration parameters used in the BirdClient
// `Path“ is the path to the BIRD unix domain socket
// `SocketBufferSize` is the size (in bytes) of the read buffer for reading from the BIRD daemon.
type BirdClientOptions struct {
	Path             string
	SocketBufferSize int
}

// New creates a new BirdClient.
//
// This function takes an optional BirdClientOptions struct which can be used to
// override the default socket path and socket buffer size.
//
// If no options are provided, the default socket path and buffer size will be used.
//
// Example:
//
//	client := gobirdc.New()
//	client := gobirdc.New(&gobirdc.BirdClientOptions{
//		Path: "/run/bird/bird.ctl",
//		SocketBufferSize: 4096,
//	}
func New(opts *BirdClientOptions) *BirdClient {
	path := "/run/bird/bird.ctl"
	socket_buffer_size := 4096

	if opts == nil {
		return &BirdClient{
			s: socket.NewBirdSocket(path, socket_buffer_size),
		}
	}

	if opts.Path != "" {
		path = opts.Path
	}
	if opts.SocketBufferSize != 0 {
		socket_buffer_size = opts.SocketBufferSize
	}

	return &BirdClient{
		s: socket.NewBirdSocket(path, socket_buffer_size),
	}
}
