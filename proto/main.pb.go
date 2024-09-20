// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/main.proto

package pb

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

type Direction int32

const (
	Direction_UP   Direction = 0
	Direction_DOWN Direction = 1
	Direction_STOP Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "UP",
		1: "DOWN",
		2: "STOP",
	}
	Direction_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
		"STOP": 2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_main_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_proto_main_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

func (x *Coordinate) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PaddleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string    `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=pb.Direction" json:"direction,omitempty"`
}

func (x *PaddleUpdateRequest) Reset() {
	*x = PaddleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaddleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaddleUpdateRequest) ProtoMessage() {}

func (x *PaddleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaddleUpdateRequest.ProtoReflect.Descriptor instead.
func (*PaddleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *PaddleUpdateRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PaddleUpdateRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_UP
}

type PaddleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PaddleUpdateResponse) Reset() {
	*x = PaddleUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaddleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaddleUpdateResponse) ProtoMessage() {}

func (x *PaddleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaddleUpdateResponse.ProtoReflect.Descriptor instead.
func (*PaddleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *PaddleUpdateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GameStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *GameStateRequest) Reset() {
	*x = GameStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStateRequest) ProtoMessage() {}

func (x *GameStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStateRequest.ProtoReflect.Descriptor instead.
func (*GameStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *GameStateRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GameStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallPos      *Coordinate `protobuf:"bytes,1,opt,name=ball_pos,json=ballPos,proto3" json:"ball_pos,omitempty"`
	P1Pos        *Coordinate `protobuf:"bytes,2,opt,name=p1_pos,json=p1Pos,proto3" json:"p1_pos,omitempty"`
	P2Pos        *Coordinate `protobuf:"bytes,3,opt,name=p2_pos,json=p2Pos,proto3" json:"p2_pos,omitempty"`
	P1Score      int32       `protobuf:"varint,4,opt,name=p1_score,json=p1Score,proto3" json:"p1_score,omitempty"`
	P2Score      int32       `protobuf:"varint,5,opt,name=p2_score,json=p2Score,proto3" json:"p2_score,omitempty"`
	ScreenHeight int32       `protobuf:"varint,6,opt,name=screen_height,json=screenHeight,proto3" json:"screen_height,omitempty"`
	ScreenWidth  int32       `protobuf:"varint,7,opt,name=screen_width,json=screenWidth,proto3" json:"screen_width,omitempty"`
}

func (x *GameStateResponse) Reset() {
	*x = GameStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStateResponse) ProtoMessage() {}

func (x *GameStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStateResponse.ProtoReflect.Descriptor instead.
func (*GameStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *GameStateResponse) GetBallPos() *Coordinate {
	if x != nil {
		return x.BallPos
	}
	return nil
}

func (x *GameStateResponse) GetP1Pos() *Coordinate {
	if x != nil {
		return x.P1Pos
	}
	return nil
}

func (x *GameStateResponse) GetP2Pos() *Coordinate {
	if x != nil {
		return x.P2Pos
	}
	return nil
}

func (x *GameStateResponse) GetP1Score() int32 {
	if x != nil {
		return x.P1Score
	}
	return 0
}

func (x *GameStateResponse) GetP2Score() int32 {
	if x != nil {
		return x.P2Score
	}
	return 0
}

func (x *GameStateResponse) GetScreenHeight() int32 {
	if x != nil {
		return x.ScreenHeight
	}
	return 0
}

func (x *GameStateResponse) GetScreenWidth() int32 {
	if x != nil {
		return x.ScreenWidth
	}
	return 0
}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x28, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79,
	0x22, 0x5f, 0x0a, 0x13, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2e, 0x0a, 0x14, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x6c,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x6c,
	0x50, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x31, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x70, 0x31, 0x50, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x32,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x05, 0x70, 0x32, 0x50, 0x6f,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x31, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x32, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x32, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x2a,
	0x27, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02,
	0x55, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x6a, 0x61, 0x6e, 0x68, 0x73,
	0x79, 0x2f, 0x70, 0x6f, 0x6e, 0x67, 0x2d, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_main_proto_goTypes = []any{
	(Direction)(0),               // 0: pb.Direction
	(*Coordinate)(nil),           // 1: pb.Coordinate
	(*PaddleUpdateRequest)(nil),  // 2: pb.PaddleUpdateRequest
	(*PaddleUpdateResponse)(nil), // 3: pb.PaddleUpdateResponse
	(*GameStateRequest)(nil),     // 4: pb.GameStateRequest
	(*GameStateResponse)(nil),    // 5: pb.GameStateResponse
}
var file_proto_main_proto_depIdxs = []int32{
	0, // 0: pb.PaddleUpdateRequest.direction:type_name -> pb.Direction
	1, // 1: pb.GameStateResponse.ball_pos:type_name -> pb.Coordinate
	1, // 2: pb.GameStateResponse.p1_pos:type_name -> pb.Coordinate
	1, // 3: pb.GameStateResponse.p2_pos:type_name -> pb.Coordinate
	2, // 4: pb.PongService.UpdatePaddleDirection:input_type -> pb.PaddleUpdateRequest
	4, // 5: pb.PongService.StreamGameState:input_type -> pb.GameStateRequest
	3, // 6: pb.PongService.UpdatePaddleDirection:output_type -> pb.PaddleUpdateResponse
	5, // 7: pb.PongService.StreamGameState:output_type -> pb.GameStateResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_main_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Coordinate); i {
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
		file_proto_main_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PaddleUpdateRequest); i {
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
		file_proto_main_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PaddleUpdateResponse); i {
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
		file_proto_main_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GameStateRequest); i {
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
		file_proto_main_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GameStateResponse); i {
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
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		EnumInfos:         file_proto_main_proto_enumTypes,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}
