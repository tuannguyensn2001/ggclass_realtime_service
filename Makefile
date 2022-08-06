gen-proto:
	rm -f src/pb/${name}/*.go
	protoc --proto_path=proto --go_out=src/pb/${name} --go_opt=paths=source_relative \
	--go-grpc_out=src/pb/${name} --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=src/pb/${name} --grpc-gateway_opt=paths=source_relative \
	proto/${name}.proto