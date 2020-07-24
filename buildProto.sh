protoc --go_out=plugins=grpc:. pkg/random/random.proto
protoc --go_out=plugins=grpc:. pkg/character/character.proto
protoc --go_out=plugins=grpc:. pkg/auth/auth.proto