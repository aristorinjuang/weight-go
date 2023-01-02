coverage:
	go test -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html
grpc:
	buf generate --template tools/grpc/buf.gen.yaml --output tools/grpc/ tools/grpc