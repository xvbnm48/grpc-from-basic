package basic

import (
	"log"

	basic "github.com/xvbnm48/grpc-basic/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Message: "Hello, World!",
	}
	// user := basic.User{
	// 	Id:       1,
	// 	Username: "xvbnm48",
	// 	IsActive: true,
	// 	Password: []byte("12345678"),
	// }

	// fmt.Println(&h)
	log.Println(&h)
	// log.Println(&user)
}
