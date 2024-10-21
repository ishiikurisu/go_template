.PHONY: default
default: test

.PHONY: test
test:
	go test dojo/*.go
