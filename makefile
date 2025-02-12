build_amd64:
	GOOS=linux GOARCH=amd64 go build -o generator_amd64 *.go