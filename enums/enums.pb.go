// The syntax for this file is proto3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: proto/enums.proto

package enums

import (
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

// we currently consider only 3 eye colours
type Student_EyeColour int32

const (
	Student_UNKNOWN_EYE_COLOUR Student_EyeColour = 0
	Student_EYE_GREEN          Student_EyeColour = 1
	Student_EYE_BROWN          Student_EyeColour = 2
	Student_EYE_BLUE           Student_EyeColour = 3
)

// Enum value maps for Student_EyeColour.
var (
	Student_EyeColour_name = map[int32]string{
		0: "UNKNOWN_EYE_COLOUR",
		1: "EYE_GREEN",
		2: "EYE_BROWN",
		3: "EYE_BLUE",
	}
	Student_EyeColour_value = map[string]int32{
		"UNKNOWN_EYE_COLOUR": 0,
		"EYE_GREEN":          1,
		"EYE_BROWN":          2,
		"EYE_BLUE":           3,
	}
)

func (x Student_EyeColour) Enum() *Student_EyeColour {
	p := new(Student_EyeColour)
	*p = x
	return p
}

func (x Student_EyeColour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Student_EyeColour) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_proto_enumTypes[0].Descriptor()
}

func (Student_EyeColour) Type() protoreflect.EnumType {
	return &file_proto_enums_proto_enumTypes[0]
}

func (x Student_EyeColour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Student_EyeColour.Descriptor instead.
func (Student_EyeColour) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_proto_rawDescGZIP(), []int{0, 0}
}

// Person is used to identify users
// across our system
type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the age as of the person's creation
	Age int32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	// the first name as documented in the signup form
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"` // last name as documented in the signup form
	// small_picture represents a small .jpg file
	SmallPicture      []byte `protobuf:"bytes,4,opt,name=small_picture,json=smallPicture,proto3" json:"small_picture,omitempty"`
	IsProfileVerified bool   `protobuf:"varint,5,opt,name=is_profile_verified,json=isProfileVerified,proto3" json:"is_profile_verified,omitempty"`
	// height of the person in cms
	Height float32 `protobuf:"fixed32,6,opt,name=height,proto3" json:"height,omitempty"`
	// a list of phone numbers that is optional to provide at signup
	PhoneNumbers []string `protobuf:"bytes,7,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	// it's an enum as defined above
	EyeColour Student_EyeColour `protobuf:"varint,8,opt,name=eye_colour,json=eyeColour,proto3,enum=enums.Student_EyeColour" json:"eye_colour,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_enums_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Student) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Student) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Student) GetSmallPicture() []byte {
	if x != nil {
		return x.SmallPicture
	}
	return nil
}

func (x *Student) GetIsProfileVerified() bool {
	if x != nil {
		return x.IsProfileVerified
	}
	return false
}

func (x *Student) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Student) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *Student) GetEyeColour() Student_EyeColour {
	if x != nil {
		return x.EyeColour
	}
	return Student_UNKNOWN_EYE_COLOUR
}

var File_proto_enums_proto protoreflect.FileDescriptor

var file_proto_enums_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xf3, 0x02, 0x0a, 0x07, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x79, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x75, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x79, 0x65, 0x43, 0x6f,
	0x6c, 0x6f, 0x75, 0x72, 0x52, 0x09, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x22,
	0x4f, 0x0a, 0x09, 0x45, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x59, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4f,
	0x55, 0x52, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x59, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x45,
	0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x59, 0x45, 0x5f, 0x42, 0x52, 0x4f, 0x57, 0x4e,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x59, 0x45, 0x5f, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x03,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_enums_proto_rawDescOnce sync.Once
	file_proto_enums_proto_rawDescData = file_proto_enums_proto_rawDesc
)

func file_proto_enums_proto_rawDescGZIP() []byte {
	file_proto_enums_proto_rawDescOnce.Do(func() {
		file_proto_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enums_proto_rawDescData)
	})
	return file_proto_enums_proto_rawDescData
}

var file_proto_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_enums_proto_goTypes = []interface{}{
	(Student_EyeColour)(0), // 0: enums.Student.EyeColour
	(*Student)(nil),        // 1: enums.Student
}
var file_proto_enums_proto_depIdxs = []int32{
	0, // 0: enums.Student.eye_colour:type_name -> enums.Student.EyeColour
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_enums_proto_init() }
func file_proto_enums_proto_init() {
	if File_proto_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
			RawDescriptor: file_proto_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enums_proto_goTypes,
		DependencyIndexes: file_proto_enums_proto_depIdxs,
		EnumInfos:         file_proto_enums_proto_enumTypes,
		MessageInfos:      file_proto_enums_proto_msgTypes,
	}.Build()
	File_proto_enums_proto = out.File
	file_proto_enums_proto_rawDesc = nil
	file_proto_enums_proto_goTypes = nil
	file_proto_enums_proto_depIdxs = nil
}