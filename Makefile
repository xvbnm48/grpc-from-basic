gen_proto-hello:
	protoc --go_out=plugins=grpc:. --go_opt=Mhello.proto=../../protogen ./proto/basic/hello.proto

gen_proto-user:
	protoc --go_out=plugins=grpc:. --go_opt=Muser.proto=../../protogen ./proto/basic/user.proto

gen_proto-userGroup:
	protoc --go_out=plugins=grpc:. --go_opt=MuserGroup.proto=../../protogen ./proto/basic/userGroup.proto

run:
	go run main.go