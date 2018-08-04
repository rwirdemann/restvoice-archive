ENVLINUX=env GOOS=linux GOARCH=amd64

linux:
	${ENVLINUX} go build ${LDFLAGS}

test:
	go test ./...

run:
	go run main.go

.PHONY: clean build linux clean test run install