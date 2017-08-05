//go:generate rm -rf data helpers server
//go:generate bash -c "go run `go env GOPATH`/src/github.com/rms1000watt/degeneres/main.go generate -f pb/main.proto -o `pwd`"

package main
