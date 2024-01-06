package basic

import (
	"log"

	"github.com/xvbnm48/grpc-basic/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	vini := basic.User{
		Id:       1,
		Username: "vini hd",
		IsActive: true,
		Password: []byte("12345678"),
		Emails:   []string{"vinihd@gmail.com"},
		Gender:   basic.Gender_FEMALE,
		Address:  &basic.Address{Street: "???", City: "???", Country: "indonesia"},
	}

	riichan := basic.User{
		Id:       2,
		Username: "riichan hd",
		IsActive: true,
		Password: []byte("12345678"),
		Emails:   []string{"riichan@gmail.com"},
		Gender:   basic.Gender_FEMALE,
		Address:  &basic.Address{Street: "??? ", City: "???", Country: "indonesia"},
	}

	userGroup := basic.UserGroup{
		GroupId:     1,
		GroupName:   "hira dazzle",
		Roles:       []string{"member", "leader"},
		Users:       []*basic.User{&vini, &riichan},
		Description: "hira dazzle is a group of people who love each other",
	}

	jsonBytes, _ := protojson.Marshal(&userGroup)
	log.Print(string(jsonBytes))
}
