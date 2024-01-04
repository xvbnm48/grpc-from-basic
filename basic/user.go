package basic

import (
	"log"

	"github.com/xvbnm48/grpc-basic/protogen/basic"
)

func BasicUser() {
	u := basic.User{
		Id:       1,
		Username: "xvbnm48",
		IsActive: true,
		Password: []byte("12345678"),
	}

	log.Print(&u)
}
