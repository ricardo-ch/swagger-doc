.PHONY: install
install:
	go get -v ./

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix cgo middleware.go