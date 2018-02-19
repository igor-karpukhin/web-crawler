SERVICE := web-crawler
PACKAGE := github.com/igor-karpukhin/${SERVICE}
COMMITMSG := ${shell git log -1 --pretty=%B}
VERSION := ${shell git describe --tags --always}
COMMIT := ${shell git rev-parse HEAD}
BUILDTIME := ${shell date -u '+%Y-%m-%d_%H:%M:%S'}
LDFLAGS := -s -w -X '${PACKAGE}/version.Version=${VERSION}' \
					-X '${PACKAGE}/version.BuildTime=${BUILDTIME}' \
					-X '${PACKAGE}/version.Commit=${COMMIT}' \
					-X '${PACKAGE}/version.CommitMsg=${COMMITMSG}'

.PHONY: clean build test all

all: clean build test

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "${LDFLAGS}" -a -o bin/${SERVICE}

build_osx:
	CGO_ENABLED=0 GOOGS=darwin go build -ldflags "${LDFLAGS}" -a -o bin/${SERVICE}

test:
	go test ./...

clean:
	rm -rf *.test
	rm -rf bin/