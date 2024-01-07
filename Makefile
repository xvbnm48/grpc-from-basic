# Change this to your own Go module
GO_MODULE := github.com/xvbnm48/grpc-basic

gen_proto-hello:
	protoc --go_out=plugins=grpc:. --go_opt=Mhello.proto=../../protogen ./proto/basic/hello.proto

gen_proto-user:
	protoc --go_out=plugins=grpc:. --go_opt=Muser.proto=../../protogen ./proto/basic/user.proto

gen_proto-userGroup:
	protoc --go_out=plugins=grpc:. --go_opt=MuserGroup.proto=../../protogen ./proto/basic/userGroup.proto

run:
	go run main.go

gen_proto-all:
	protoc --go_out=plugins=grpc:. --go_opt=Mhello.proto=../../protogen ./proto/basic/hello.proto
	protoc --go_out=plugins=grpc:. --go_opt=Muser.proto=../../protogen ./proto/basic/user.proto
	protoc --go_out=plugins=grpc:. --go_opt=MuserGroup.proto=../../protogen ./proto/basic/userGroup.proto
	protoc --go_out=plugins=grpc:. --go_opt=Mapplication.proto=../../protogen ./proto/dummy/application.proto
	protoc --go_out=plugins=grpc:. --go_opt=Mapplication.proto=../../protogen ./proto/basic/application.proto
	protoc --go_out=plugins=grpc:. --go_opt=Mjobsearch.proto=../../protogen ./proto/jobsearch/jobsearch.proto

gen_proto-all2:
	protoc --go_out=plugins=grpc:. --go_opt=Mhello.proto=../../protogen ./proto/basic/hello.proto \
	       --go_out=plugins=grpc:. --go_opt=Muser.proto=../../protogen ./proto/basic/user.proto \
	       --go_out=plugins=grpc:. --go_opt=MuserGroup.proto=../../protogen ./proto/basic/userGroup.proto \
	       --go_out=plugins=grpc:. --go_opt=Mapplication.proto=../../protogen ./proto/dummy/application.proto \
	       --go_out=plugins=grpc:. --go_opt=Mapplication.proto=../../protogen ./proto/basic/application.proto \
	       --go_out=plugins=grpc:. --go_opt=Mjobsearch.proto=../../protogen ./proto/jobsearch/jobsearch.proto
