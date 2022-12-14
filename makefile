build:
	go build -ldflags="-s -w" -o bin/ms-mvc main.go

run:
	go run main.go

tests:
	@read -p "Enter the directory you wish to run tests from:" directory; \
	test_dir=$$directory; \
	go test -cover ./$$test_dir -v

compile:
	echo "Compiling ms-mvc app for each OS/Platform/architecture"
	GOOS="darwin" GOARCH="arm64" go build -ldflags="-s -w" -o bin/ms-mvc-osx-arm64 main.go
	GOOS="darwin" GOARCH="amd64" go build -ldflags="-s -w" -o bin/ms-mvc-osx-amd64 main.go
	GOOS="windows" GOARCH="386" go build -ldflags="-s -w" -o bin/ms-mvc-windows-386.exe main.go
	GOOS="windows" GOARCH="amd64" go build -ldflags="-s -w" -o bin/ms-mvc-windows-amd64.exe main.go
	GOOS="linux" GOARCH="amd64" go build -ldflags="-s -w" -o bin/ms-mvc-linux-amd64 main.go
	GOOS="linux" GOARCH="386" go build -ldflags="-s -w" -o bin/ms-mvc-linux-386 main.go
