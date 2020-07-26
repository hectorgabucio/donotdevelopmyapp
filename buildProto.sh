protoc --go_out=plugins=grpc:. internal/random/random.proto
protoc --go_out=plugins=grpc:. internal/character/character.proto
protoc --go_out=plugins=grpc:. internal/auth/auth.proto