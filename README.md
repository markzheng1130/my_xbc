# my_xbc

### [install]
- https://grpc.io/docs/languages/go/quickstart/
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

### [init go mod]
- step 1, cd to "my_xbc" folder
- step 2, go mod init mx
- step 3, go mod tidy

### [run]
- step 1, go run ./mock_ams_server/main.go
- step 2, go run ./mock_grpc_api_client/main.go

### [re-build proto]
- step 1, cd to "my_xbc" folder
- step 2, protoc --go_out=. --go-grpc_out=. outbox_service/*.proto
