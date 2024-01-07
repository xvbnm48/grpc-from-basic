package jobsearch

import (
	"encoding/json"
	"log"

	"github.com/xvbnm48/grpc-basic/protogen/basic"
	"github.com/xvbnm48/grpc-basic/protogen/dummy"
	"github.com/xvbnm48/grpc-basic/protogen/jobsearch"
	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftaware{
		JobSoftwareId: 1,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "jobsearch",
			Platforms: []string{"linux", "windows", "macos"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("software:", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandiate{
		JobCandidateId: 1,
		Application: &dummy.Application{
			ApplicationId:       1,
			ApplicationFullName: "vini hd",
			Email:               "vinihd@gmail.com",
			Phone:               "08123456789",
		},
	}
	jsonBytes, _ := json.Marshal(&jc)
	log.Println("candidate:", string(jsonBytes))
}
