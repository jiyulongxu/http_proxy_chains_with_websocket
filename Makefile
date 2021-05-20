all:
	go build -o ./bin/client ./cmd/client.go
	go build -o ./bin/server ./cmd/server.go

windows:
	CGO_ENABLED=0 GOOS=windows go build -o ./bin/client.exe ./cmd/client.go
	CGO_ENABLED=0 GOOS=windows go build -o ./bin/server.exe ./cmd/server.go

mac:
	CGO_ENABLED=0 GOOS=darwin go build -o ./bin/client ./cmd/client.go
	CGO_ENABLED=0 GOOS=darwin go build -o ./bin/server ./cmd/server.go

linux_mips:
	CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o ./bin/client ./cmd/client.go
	CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o ./bin/server ./cmd/server.go

linux_armv5:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/client ./cmd/client.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/server ./cmd/server.go

linux_armv7:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o ./bin/client ./cmd/client.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o ./bin/server ./cmd/server.go

openssl_key:
	openssl genrsa -out ./bin/server.key 2048
	openssl req -nodes -new -key ./bin/server.key -subj "/CN=localhost" -out ./bin/server.csr
	openssl x509 -req -sha256 -days 365 -in ./bin/server.csr -signkey ./bin/server.key -out ./bin/server.crt

clean:
	rm -rfd ./bin/*
