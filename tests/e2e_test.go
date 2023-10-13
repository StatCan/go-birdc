/*
For now these tests are designed to be run in the supplied container (built with Dockerfile)
in a devcontainer environment.

A similar environment can be provisioned on a linux server by simply installing go1.20+ and Bird2.

At the present time, I do not have the e2e tests below configured to work in gitlab-ci.

TODO: implement e2e tests in gitlab-ci
*/
package e2e_test

import (
	"regexp"
	"strings"
	"testing"

	gobirdc "go.sykesdev.ca/go-birdc"
)

const TEST_BIRD_SOCKET_PATH = "/usr/local/var/run/bird.ctl"

var completeCodeRegexp *regexp.Regexp

func init() {
	completeCodeRegexp = regexp.MustCompile(`(?m)^([089][0-9]{3})`)
}

func TestShowStatus(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, replyCode, err := b.ShowStatus()
	if err != nil {
		t.Fatalf("got an error running 'ShowStatus'. %s. %s", err, replyCode)
	}
	t.Log(string(resp))

	if string(replyCode) != "0013" {
		t.Fatalf("did not get the expected outcome. %s. %s.", replyCode, resp)
	}
}

func TestShowMemory(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowMemory()
	if err != nil {
		t.Fatalf("got an error running 'ShowMemory'. %s", err)
	}
	t.Log(string(resp))

	if !strings.Contains(string(resp), "1018-BIRD memory usage") {
		t.Fatalf("did not get the expected outcome. %s", resp)
	}
}

func TestShowProtocols(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowProtocols()
	if err != nil {
		t.Fatalf("got an error running 'ShowProtocols'. %s", err)
	}
	t.Log(string(resp))

	if !strings.Contains(string(resp), "2002-Name       Proto") {
		t.Fatalf("did not get the expected outcome. %s", resp)
	}
}

func TestShowInterfaces(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowInterfaces()
	if err != nil {
		t.Fatalf("got an error running 'ShowInterfaces'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowRoute(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowRoute()
	if err != nil {
		t.Fatalf("got an error running 'ShowRoute'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowSymbols(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowSymbols()
	if err != nil {
		t.Fatalf("got an error running 'ShowSymbols'. %s", err)
	}
	t.Log(string(resp))

	if !strings.Contains(string(resp), "routing table") && !strings.Contains(string(resp), "protocol") {
		t.Fatalf("did not get the expected outcome. %s", resp)
	}
}

func TestShowBFDSessions(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowBFDSessions()
	if err != nil {
		t.Fatalf("got an error running 'ShowBFDSessions'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowBabelInterfaces(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowBabelInterfaces()
	if err != nil {
		t.Fatalf("got an error running 'ShowBabelInterfaces'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowBabelNeighbors(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowBabelNeighbors()
	if err != nil {
		t.Fatalf("got an error running 'ShowBabelNeighbors'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowBabelEntries(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowBabelEntries()
	if err != nil {
		t.Fatalf("got an error running 'ShowBabelEntries'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowBabelRoutes(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowBabelRoutes()
	if err != nil {
		t.Fatalf("got an error running 'ShowBabelRoutes'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowOSPF(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowOSPF()
	if err != nil {
		t.Fatalf("got an error running 'ShowOSPF'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowRIPInterfaces(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowRIPInterfaces()
	if err != nil {
		t.Fatalf("got an error running 'ShowRIPInterfaces'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowRIPNeighbors(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowRIPNeighbors()
	if err != nil {
		t.Fatalf("got an error running 'ShowRIPNeighbors'. %s", err)
	}
	t.Log(string(resp))
}

func TestShowStatic(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.ShowStatic()
	if err != nil {
		t.Fatalf("got an error running 'ShowStatic'. %s", err)
	}
	t.Log(string(resp))
}

func TestConfigure(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.Configure()
	if err != nil {
		t.Fatalf("got an error running 'Configure'. %s", err)
	}
	t.Log(string(resp))
}

func TestConfigureSoft(t *testing.T) {
	b := gobirdc.New(TEST_BIRD_SOCKET_PATH)

	resp, _, err := b.Configure("soft")
	if err != nil {
		t.Fatalf("got an error running 'Configure'. %s", err)
	}
	t.Log(string(resp))
}
