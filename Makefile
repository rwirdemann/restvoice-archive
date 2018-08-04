ENVLINUX=env GOOS=linux GOARCH=amd64

linux:
	${ENVLINUX} go build ${LDFLAGS}

test:
	go test ./...

run:
	go run main.go

deploy: linux
	scp keycloak_key.pub restvoice@178.128.203.76:~
	ssh restvoice@178.128.203.76 "nohup ./restvoice &"

.PHONY: deploy linux test run