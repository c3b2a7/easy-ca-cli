NAME=easy-ca-cli
BINDIR=bin
VERSION=$(shell git describe --dirty --always --tags || echo "unknown version")
COMMIT_HASH=$(shell git rev-parse HEAD)
BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags ' \
		-X "github.com/c3b2a7/easy-ca-cli/cmd.version=$(VERSION)" \
		-X "github.com/c3b2a7/easy-ca-cli/cmd.date=$(BUILDTIME)" \
		-X "github.com/c3b2a7/easy-ca-cli/cmd.commit=$(COMMIT_HASH)" \
		-X "github.com/c3b2a7/easy-ca-cli/cmd.builtBy=make" \
		-w -s -buildid='

all: linux-amd64 linux-arm64 freebsd-amd64 freebsd-arm64 macos-amd64 macos-arm64 win64 win32

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

freebsd-arm64:
	GOARCH=arm64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

macos-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

macos-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

win64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

win32:
	GOARCH=386 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

test: test-linux-amd64 test-linux-arm64 test-freebsd-amd64 test-freebsd-arm64 test-macos-amd64 test-macos-arm64 test-win64 test-win32

test-linux-amd64:
	GOARCH=amd64 GOOS=linux go test

test-linux-arm64:
	GOARCH=arm64 GOOS=linux go test

test-freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd go test

test-freebsd-arm64:
	GOARCH=arm64 GOOS=freebsd go test

test-macos-amd64:
	GOARCH=amd64 GOOS=darwin go test

test-macos-arm64:
	GOARCH=arm64 GOOS=darwin go test

test-win64:
	GOARCH=amd64 GOOS=windows go test

test-win32:
	GOARCH=386 GOOS=windows go test

clean:
	rm $(BINDIR)/*