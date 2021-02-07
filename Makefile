generate:
	buf generate --path ./proto/common.proto
	buf generate --path ./proto/tag.proto
	buf generate --path ./proto/address.proto
	buf generate --path ./proto/streetaddress.proto
	buf generate --path ./proto/user.proto
	buf generate --path ./proto/usergroup.proto
	# Generate static assets for OpenAPI UI
	statik -m -f -src third_party/OpenAPI/

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf
