package basic

import (
	"log"

	"github.com/xvbnm48/grpc-basic/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "pondok cabe tangerang selatan",
		City:    "tangerang",
		Country: "indonesia",
	}
	u := basic.User{
		Id:       1,
		Username: "xvbnm48",
		IsActive: true,
		Password: []byte("12345678"),
		Gender:   basic.Gender_MALE,
		Emails:   []string{"xvbnm48@duckduckgo.com", "muhamadfarizwisnuprananda@gmail.com"},
		Address:  &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)

	log.Print(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       2,
		Username: "vini",
		IsActive: true,
		Password: []byte("12345678"),
		Emails:   []string{"vinihd@gmail.com"},
		Gender:   basic.Gender_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Print(string(jsonBytes))
}

func JsonToProto() {
	json := `{"id":2, "username":"vini", "is_active":true, "password":"MTIzNDU2Nzg=", "emails":["vinihd@gmail.com"], "gender":"FEMALE"}`
	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatal("error unmarshaling json: ", err)
	}

	log.Print(&p)
}
