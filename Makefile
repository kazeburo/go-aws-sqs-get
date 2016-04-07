VERSION=0.0.1

all: aws-sqs-get

.PHONY: aws-sqs-get

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

aws-sqs-get: aws-sqs-get.go
	gom build -o aws-sqs-get

linux: aws-sqs-get.go
	GOOS=linux GOARCH=amd64 gom build -o aws-sqs-get

fmt:
	go fmt ./...

dist:
	git archive --format tgz HEAD -o aws-sqs-get-$(VERSION).tar.gz --prefix aws-sqs-get-$(VERSION)/

clean:
	rm -rf aws-sqs-get aws-sqs-get-*.tar.gz

