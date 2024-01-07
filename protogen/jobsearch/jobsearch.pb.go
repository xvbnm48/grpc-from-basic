// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: proto/jobsearch/jobsearch.proto

package jobsearch

import (
	basic "github.com/xvbnm48/grpc-basic/protogen/basic"
	dummy "github.com/xvbnm48/grpc-basic/protogen/dummy"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobCandiate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCandidateId uint32             `protobuf:"varint,1,opt,name=job_candidate_id,json=jobCandidate_id,proto3" json:"job_candidate_id,omitempty"`
	Application    *dummy.Application `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *JobCandiate) Reset() {
	*x = JobCandiate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobCandiate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCandiate) ProtoMessage() {}

func (x *JobCandiate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCandiate.ProtoReflect.Descriptor instead.
func (*JobCandiate) Descriptor() ([]byte, []int) {
	return file_proto_jobsearch_jobsearch_proto_rawDescGZIP(), []int{0}
}

func (x *JobCandiate) GetJobCandidateId() uint32 {
	if x != nil {
		return x.JobCandidateId
	}
	return 0
}

func (x *JobCandiate) GetApplication() *dummy.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

type JobSoftaware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobSoftwareId uint32             `protobuf:"varint,1,opt,name=job_software_id,json=jobSoftware_id,proto3" json:"job_software_id,omitempty"`
	Application   *basic.Application `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *JobSoftaware) Reset() {
	*x = JobSoftaware{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSoftaware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSoftaware) ProtoMessage() {}

func (x *JobSoftaware) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSoftaware.ProtoReflect.Descriptor instead.
func (*JobSoftaware) Descriptor() ([]byte, []int) {
	return file_proto_jobsearch_jobsearch_proto_rawDescGZIP(), []int{1}
}

func (x *JobSoftaware) GetJobSoftwareId() uint32 {
	if x != nil {
		return x.JobSoftwareId
	}
	return 0
}

func (x *JobSoftaware) GetApplication() *basic.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_proto_jobsearch_jobsearch_proto protoreflect.FileDescriptor

var file_proto_jobsearch_jobsearch_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x72, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x61, 0x74, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x74, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x53, 0x6f, 0x66, 0x74, 0x61, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x77,
	0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6a, 0x6f,
	0x62, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68, 0x65, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x76, 0x62, 0x6e, 0x6d, 0x34, 0x38, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_jobsearch_jobsearch_proto_rawDescOnce sync.Once
	file_proto_jobsearch_jobsearch_proto_rawDescData = file_proto_jobsearch_jobsearch_proto_rawDesc
)

func file_proto_jobsearch_jobsearch_proto_rawDescGZIP() []byte {
	file_proto_jobsearch_jobsearch_proto_rawDescOnce.Do(func() {
		file_proto_jobsearch_jobsearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_jobsearch_jobsearch_proto_rawDescData)
	})
	return file_proto_jobsearch_jobsearch_proto_rawDescData
}

var file_proto_jobsearch_jobsearch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_jobsearch_jobsearch_proto_goTypes = []interface{}{
	(*JobCandiate)(nil),       // 0: JobCandiate
	(*JobSoftaware)(nil),      // 1: JobSoftaware
	(*dummy.Application)(nil), // 2: the.dummy.Application
	(*basic.Application)(nil), // 3: the.basic.Application
}
var file_proto_jobsearch_jobsearch_proto_depIdxs = []int32{
	2, // 0: JobCandiate.application:type_name -> the.dummy.Application
	3, // 1: JobSoftaware.application:type_name -> the.basic.Application
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_jobsearch_jobsearch_proto_init() }
func file_proto_jobsearch_jobsearch_proto_init() {
	if File_proto_jobsearch_jobsearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_jobsearch_jobsearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobCandiate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_jobsearch_jobsearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSoftaware); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_jobsearch_jobsearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_jobsearch_jobsearch_proto_goTypes,
		DependencyIndexes: file_proto_jobsearch_jobsearch_proto_depIdxs,
		MessageInfos:      file_proto_jobsearch_jobsearch_proto_msgTypes,
	}.Build()
	File_proto_jobsearch_jobsearch_proto = out.File
	file_proto_jobsearch_jobsearch_proto_rawDesc = nil
	file_proto_jobsearch_jobsearch_proto_goTypes = nil
	file_proto_jobsearch_jobsearch_proto_depIdxs = nil
}
