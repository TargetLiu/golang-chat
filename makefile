build: 
	GO111MODULE=on go build -o ./bin/golang-chat ./cmd/golang-chat

install: 
	GO111MODULE=on go install ./cmd/golang-chat