VERSION = 0.0.2-SNAPSHOT
LDFLAGS = -s -w -X main.version=${VERSION}

sources := clusterhat.go

all: compile

deps:
	go get github.com/warthog618/gpio
	go get github.com/alexflint/go-arg

compile: $(sources) deps
	GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="${LDFLAGS}"
