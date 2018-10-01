ENVLINUX=env GOOS=linux GOARCH=amd64
VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.githash=${VERSION} -X main.buildstamp=${BUILD}"

linux:
	${ENVLINUX} go build ${LDFLAGS}

test:
	go test ./...

run:
	go run main.go

deploy: linux
	ssh restvoice@restvoice.org "killall -q restvoice || true"
	scp keycloak_key.pub restvoice@restvoice.org:~
	scp restvoice restvoice@restvoice.org:~
	ssh restvoice@restvoice.org "nohup ./restvoice > restvoice.out 2> restvoice.err < /dev/null &"
	curl http://restvoice.org:8080/version

.PHONY: deploy linux test run