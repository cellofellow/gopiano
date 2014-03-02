# Pandora.com API Wrapper Client

A very thin wrapper around Pandora.com's JSON API.

## Usage and Hacking

First create or set a `$GOPATH`, then use `go get`.

```sh
go get github.com/cellofellow/gopiano
```

You can then import this into your own code with

```go
import "github.com/cellofellow/gopiano"
```

Or if you like you hack on it `cd $GOPATH/src/github.com/cellofellow/gopiano`.
I'm very much in need of someone with Go experience to vet my code, and some
specific things still need doing:

* Proper tests.
* Proper error handling.

This is *alpha quality code*, use at your own risk.
