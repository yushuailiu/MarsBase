generate:
	go-assets-builder config -o config/config.go -p config
linux:
	GOOS=linux GOARCH=amd64  go build -o MarsBase
windows:
	GOOS=windows GOARCH=amd64  go build -o MarsBase
mac:
	GOOS=darwin GOARCH=amd64  go build -o MarsBase