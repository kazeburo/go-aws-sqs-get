VERSION=0.0.2
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
GO111MODULE=on

all: aws-sqs-get

.PHONY: aws-sqs-get

aws-sqs-get: aws-sqs-get.go
	go build $(LDFLAGS) -o aws-sqs-get

linux: aws-sqs-get.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o aws-sqs-get

clean:
	rm -rf aws-sqs-get

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
	goreleaser --rm-dist
