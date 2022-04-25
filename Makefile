generate: 
	buf generate --path ./proto/user/common.proto
	buf generate --path ./proto/user/tag.proto
	buf generate --path ./proto/user/address.proto
	buf generate --path ./proto/user/streetaddress.proto
	buf generate --path ./proto/user/user_messages.proto
	buf generate --path ./proto/user/usergroup_messages.proto
	buf generate --path ./proto/user/user.proto
	# Generate static assets for OpenAPI UI
	statik -m -f -src third_party/OpenAPI/

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf
