.PHONY: tools lint fmt test build tidy release

NAME=go-template-project
BINDIR=dist

# -- begin inject some build information --
PACKAGE=github.com/xpzouying/go-cmd-project-template/internal/constant
BUILDTIME=$(shell date +"%Y-%m-%d-%T")
BRANCH=$(shell git rev-parse --abbrev-ref HEAD | tr -d '\040\011\012\015\n')
REVISION=$(shell git rev-parse HEAD)
VERSION="0.0.1"
# -- end inject some build information --


# -w -s：You will get the smallest binaries if you compile with -ldflags '-w -s'. The -w turns off DWARF debugging information
# for more information, please refer to https://stackoverflow.com/questions/22267189/what-does-the-w-flag-mean-when-passed-in-via-the-ldflags-option-to-the-go-comman
# we need CGO_ENABLED=1 because we import the node_exporter ,and we need install `glibc-source,libc6` to make it work
# TODO check if node_exporter collector with CGO_ENABLED=0 is enough
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags="-w -s -X ${PACKAGE}.GitBranch=${BRANCH} -X ${PACKAGE}.GitRevision=${REVISION} -X ${PACKAGE}.BuildTime=${BUILDTIME} -X ${PACKAGE}.Version=${VERSION}"

test:
	go test -v -count=1 -timeout=1m ./...

build:
	${GOBUILD} -o $(BINDIR)/$(NAME) cmd/main.go

build-arm:
	GOARCH=arm GOOS=linux ${GOBUILD} -o $(BINDIR)/$(NAME) cmd/ehco/main.go

build-linux-amd64:
	GOARCH=amd64 GOOS=linux ${GOBUILD} -o $(BINDIR)/$(NAME)_amd64 cmd/ehco/main.go

tidy:
	go mod tidy