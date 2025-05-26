MAIN:=./cmd/server
TARGET:=location-service
WIN_TARGET:=${TARGET}.exe
SERVER_CMD:=./${TARGET}
PROTOC_DEST:=./
PROTOC_FLAGS:=--go_out=${PROTOC_DEST} --go_opt=paths=source_relative --go-grpc_out=${PROTOC_DEST} --go-grpc_opt=paths=source_relative
USERS_PROTO_FILES:=./grpc/users/users.proto
POSTS_PROTO_FILES:=./grpc/posts/posts.proto
RESTAURANTS_PROTO_FILES:=./grpc/restaurants/restaurants.proto
build-proto:
	protoc ${PROTOC_FLAGS} ${POSTS_PROTO_FILES}
