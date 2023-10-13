package gobirdc

import (
	"sykesdev.ca/go-birdc/socket"
)

type BirdClient struct {
	s *socket.BirdSocket
}

func New(path ...string) *BirdClient {
	if len(path) == 0 {
		path = []string{"/run/bird/bird.ctl"}
	}
	return &BirdClient{
		s: socket.NewBirdSocket(path[0], 4096),
	}
}
