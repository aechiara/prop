EXEC_NAME = prop

all:
	go build -o $(EXEC_NAME)

test:
	go test ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o $(EXEC_NAME)-linux

windows:
	set GOOS=windows
	set GOARCH=amd64
	go build -o $(EXEC_NAME).exe

run:
	./${EXEC_NAME}


clean:
	rm -rf $(EXEC_NAME) $(EXEC_NAME)-linux $(EXEC_NAME).exe