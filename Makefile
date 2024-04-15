.PHONY: test
test:
	go test -cover ./...

t="coverage.txt"
coverage:
	go test -coverprofile=$t ./... && go tool cover -html=$t && unlink $t


