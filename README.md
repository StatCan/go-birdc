# go-birdc

A small GO library for working with the BIRD internet routing daemon through the provided unix socket.

## Usage

Install the library as a dependency
```bash
# Directly from source
go get -u https://gitlab.sykesdev.ca/standalone-projects/go-birdc

# Through goproxy
go get -u sykesdev.ca/go-birdc@latest
```

Then start using it

```golang
import "sykesdev.ca/go-birdc"

func main() {
	b := New() // optionally you can specify a path to the unix socket endpoint (defaults to: /run/bird/bird.ctl)
	resp, replyCode, err := b.ShowStatus()
	if err != nil {
		log.Fatalf("something went wrong. %s", replyCode)
	}

	resp, replyCode, err := b.RestartProtocol("bgp")
	if err != nil {
		log.Fatalf("somthing went wrong. %s", replyCode)
	}

	fmt.Println(string(resp))
}

```

And that's it!
