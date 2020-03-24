EXEC_NAME = prop
EXEC_VERSION = $(shell grep "const VERSION" cmd/version.go | cut -d" " -f4)

all:
	go build -o $(EXEC_NAME)

dist: clean
	mkdir dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(EXEC_NAME)-linux
	GOOS=darwin GOARCH=amd64 go build -o dist/$(EXEC_NAME)-macos
	set GOOS=windows
	set GOARCH=amd64
	go build -o dist/$(EXEC_NAME).exe
	cd dist && zip prop-linux-$(EXEC_VERSION).zip $(EXEC_NAME)-linux
	cd dist && zip prop-macos-$(EXEC_VERSION).zip $(EXEC_NAME)-macos
	cd dist && zip prop-windows-$(EXEC_VERSION).zip $(EXEC_NAME).exe

test:
	go test ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o $(EXEC_NAME)-linux

macos:
	GOOS=darwin GOARCH=amd64 go build -o $(EXEC_NAME)-macos

windows:
	set GOOS=windows
	set GOARCH=amd64
	go build -o $(EXEC_NAME).exe

run:
	./${EXEC_NAME}


clean:
	rm -rf $(EXEC_NAME) $(EXEC_NAME)-linux $(EXEC_NAME).exe $(EXEC_NAME)-macos
	rm -rf dist
	