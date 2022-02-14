BIN := sentry-go-test

build:
	CGO_ENABLED=0 go build -o $(BIN) -ldflags "-s -w" .

clean:
	-rm -f $(BIN)
