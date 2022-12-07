.PHONY: build
build:
	go build -o gostudy cmd/gostudy/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -f gostudy
