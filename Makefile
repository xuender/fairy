PACKAGE = github.com/xuender/fairy

tools:
	go install fyne.io/fyne/v2/cmd/fyne@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/spf13/cobra-cli@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/cespare/reflex@latest

test:
	go test -race -v ./... -gcflags=all=-l

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./...'

clean:
	rm -rf dist

build:
	go build -o dist/fairy main.go

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

proto:
	protoc --go_out=. pb/*.proto

wire:
	wire gen ${PACKAGE}/cmd
