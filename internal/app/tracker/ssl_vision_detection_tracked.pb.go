// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: tracker/ssl_vision_detection_tracked.proto

package tracker

import (
	geom "github.com/RoboCup-SSL/ssl-game-controller/internal/app/geom"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
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

// Capabilities that a source implementation can have
type Capability int32

const (
	Capability_CAPABILITY_UNKNOWN               Capability = 0
	Capability_CAPABILITY_DETECT_FLYING_BALLS   Capability = 1
	Capability_CAPABILITY_DETECT_MULTIPLE_BALLS Capability = 2
	Capability_CAPABILITY_DETECT_KICKED_BALLS   Capability = 3
)

// Enum value maps for Capability.
var (
	Capability_name = map[int32]string{
		0: "CAPABILITY_UNKNOWN",
		1: "CAPABILITY_DETECT_FLYING_BALLS",
		2: "CAPABILITY_DETECT_MULTIPLE_BALLS",
		3: "CAPABILITY_DETECT_KICKED_BALLS",
	}
	Capability_value = map[string]int32{
		"CAPABILITY_UNKNOWN":               0,
		"CAPABILITY_DETECT_FLYING_BALLS":   1,
		"CAPABILITY_DETECT_MULTIPLE_BALLS": 2,
		"CAPABILITY_DETECT_KICKED_BALLS":   3,
	}
)

func (x Capability) Enum() *Capability {
	p := new(Capability)
	*p = x
	return p
}

func (x Capability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Capability) Descriptor() protoreflect.EnumDescriptor {
	return file_tracker_ssl_vision_detection_tracked_proto_enumTypes[0].Descriptor()
}

func (Capability) Type() protoreflect.EnumType {
	return &file_tracker_ssl_vision_detection_tracked_proto_enumTypes[0]
}

func (x Capability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Capability) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Capability(num)
	return nil
}

// Deprecated: Use Capability.Descriptor instead.
func (Capability) EnumDescriptor() ([]byte, []int) {
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP(), []int{0}
}

// A single tracked ball
type TrackedBall struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The position (x, y, height) [m] in the ssl-vision coordinate system
	Pos *geom.Vector3 `protobuf:"bytes,1,req,name=pos" json:"pos,omitempty"`
	// The velocity [m/s] in the ssl-vision coordinate system
	Vel *geom.Vector3 `protobuf:"bytes,2,opt,name=vel" json:"vel,omitempty"`
	// The visibility of the ball
	// A value between 0 (not visible) and 1 (visible)
	// The exact implementation depends on the source software
	Visibility    *float32 `protobuf:"fixed32,3,opt,name=visibility" json:"visibility,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackedBall) Reset() {
	*x = TrackedBall{}
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedBall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedBall) ProtoMessage() {}

func (x *TrackedBall) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedBall.ProtoReflect.Descriptor instead.
func (*TrackedBall) Descriptor() ([]byte, []int) {
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP(), []int{0}
}

func (x *TrackedBall) GetPos() *geom.Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *TrackedBall) GetVel() *geom.Vector3 {
	if x != nil {
		return x.Vel
	}
	return nil
}

func (x *TrackedBall) GetVisibility() float32 {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return 0
}

// A ball kicked by a robot, including predictions when the ball will come to a stop
type KickedBall struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The initial position [m] from which the ball was kicked
	Pos *geom.Vector2 `protobuf:"bytes,1,req,name=pos" json:"pos,omitempty"`
	// The initial velocity [m/s] with which the ball was kicked
	Vel *geom.Vector3 `protobuf:"bytes,2,req,name=vel" json:"vel,omitempty"`
	// The unix timestamp [s] when the kick was performed
	StartTimestamp *float64 `protobuf:"fixed64,3,req,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	// The predicted unix timestamp [s] when the ball comes to a stop
	StopTimestamp *float64 `protobuf:"fixed64,4,opt,name=stop_timestamp,json=stopTimestamp" json:"stop_timestamp,omitempty"`
	// The predicted position [m] at which the ball will come to a stop
	StopPos *geom.Vector2 `protobuf:"bytes,5,opt,name=stop_pos,json=stopPos" json:"stop_pos,omitempty"`
	// The robot that kicked the ball
	RobotId       *state.RobotId `protobuf:"bytes,6,opt,name=robot_id,json=robotId" json:"robot_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KickedBall) Reset() {
	*x = KickedBall{}
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KickedBall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickedBall) ProtoMessage() {}

func (x *KickedBall) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickedBall.ProtoReflect.Descriptor instead.
func (*KickedBall) Descriptor() ([]byte, []int) {
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP(), []int{1}
}

func (x *KickedBall) GetPos() *geom.Vector2 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *KickedBall) GetVel() *geom.Vector3 {
	if x != nil {
		return x.Vel
	}
	return nil
}

func (x *KickedBall) GetStartTimestamp() float64 {
	if x != nil && x.StartTimestamp != nil {
		return *x.StartTimestamp
	}
	return 0
}

func (x *KickedBall) GetStopTimestamp() float64 {
	if x != nil && x.StopTimestamp != nil {
		return *x.StopTimestamp
	}
	return 0
}

func (x *KickedBall) GetStopPos() *geom.Vector2 {
	if x != nil {
		return x.StopPos
	}
	return nil
}

func (x *KickedBall) GetRobotId() *state.RobotId {
	if x != nil {
		return x.RobotId
	}
	return nil
}

// A single tracked robot
type TrackedRobot struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	RobotId *state.RobotId         `protobuf:"bytes,1,req,name=robot_id,json=robotId" json:"robot_id,omitempty"`
	// The position [m] in the ssl-vision coordinate system
	Pos *geom.Vector2 `protobuf:"bytes,2,req,name=pos" json:"pos,omitempty"`
	// The orientation [rad] in the ssl-vision coordinate system
	Orientation *float32 `protobuf:"fixed32,3,req,name=orientation" json:"orientation,omitempty"`
	// The velocity [m/s] in the ssl-vision coordinate system
	Vel *geom.Vector2 `protobuf:"bytes,4,opt,name=vel" json:"vel,omitempty"`
	// The angular velocity [rad/s] in the ssl-vision coordinate system
	VelAngular *float32 `protobuf:"fixed32,5,opt,name=vel_angular,json=velAngular" json:"vel_angular,omitempty"`
	// The visibility of the robot
	// A value between 0 (not visible) and 1 (visible)
	// The exact implementation depends on the source software
	Visibility    *float32 `protobuf:"fixed32,6,opt,name=visibility" json:"visibility,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackedRobot) Reset() {
	*x = TrackedRobot{}
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedRobot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedRobot) ProtoMessage() {}

func (x *TrackedRobot) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedRobot.ProtoReflect.Descriptor instead.
func (*TrackedRobot) Descriptor() ([]byte, []int) {
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP(), []int{2}
}

func (x *TrackedRobot) GetRobotId() *state.RobotId {
	if x != nil {
		return x.RobotId
	}
	return nil
}

func (x *TrackedRobot) GetPos() *geom.Vector2 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *TrackedRobot) GetOrientation() float32 {
	if x != nil && x.Orientation != nil {
		return *x.Orientation
	}
	return 0
}

func (x *TrackedRobot) GetVel() *geom.Vector2 {
	if x != nil {
		return x.Vel
	}
	return nil
}

func (x *TrackedRobot) GetVelAngular() float32 {
	if x != nil && x.VelAngular != nil {
		return *x.VelAngular
	}
	return 0
}

func (x *TrackedRobot) GetVisibility() float32 {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return 0
}

// A frame that contains all currently tracked objects on the field on all cameras
type TrackedFrame struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A monotonous increasing frame counter
	FrameNumber *uint32 `protobuf:"varint,1,req,name=frame_number,json=frameNumber" json:"frame_number,omitempty"`
	// The unix timestamp in [s] of the data
	// If timestamp is larger than timestamp_captured, the source has applied a prediction already
	Timestamp *float64 `protobuf:"fixed64,2,req,name=timestamp" json:"timestamp,omitempty"`
	// The list of detected balls
	// The first ball is the primary one
	// Sources may add additional balls based on their capabilities
	Balls []*TrackedBall `protobuf:"bytes,3,rep,name=balls" json:"balls,omitempty"`
	// The list of detected robots of both teams
	Robots []*TrackedRobot `protobuf:"bytes,4,rep,name=robots" json:"robots,omitempty"`
	// Information about a kicked ball, if the ball was kicked by a robot and is still moving
	// Note: This field is optional. Some source implementations might not set this at any time
	KickedBall *KickedBall `protobuf:"bytes,5,opt,name=kicked_ball,json=kickedBall" json:"kicked_ball,omitempty"`
	// List of capabilities of the source implementation
	Capabilities  []Capability `protobuf:"varint,6,rep,name=capabilities,enum=Capability" json:"capabilities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackedFrame) Reset() {
	*x = TrackedFrame{}
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedFrame) ProtoMessage() {}

func (x *TrackedFrame) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_ssl_vision_detection_tracked_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedFrame.ProtoReflect.Descriptor instead.
func (*TrackedFrame) Descriptor() ([]byte, []int) {
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP(), []int{3}
}

func (x *TrackedFrame) GetFrameNumber() uint32 {
	if x != nil && x.FrameNumber != nil {
		return *x.FrameNumber
	}
	return 0
}

func (x *TrackedFrame) GetTimestamp() float64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *TrackedFrame) GetBalls() []*TrackedBall {
	if x != nil {
		return x.Balls
	}
	return nil
}

func (x *TrackedFrame) GetRobots() []*TrackedRobot {
	if x != nil {
		return x.Robots
	}
	return nil
}

func (x *TrackedFrame) GetKickedBall() *KickedBall {
	if x != nil {
		return x.KickedBall
	}
	return nil
}

func (x *TrackedFrame) GetCapabilities() []Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

var File_tracker_ssl_vision_detection_tracked_proto protoreflect.FileDescriptor

var file_tracker_ssl_vision_detection_tracked_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x65, 0x6f, 0x6d, 0x2f, 0x73, 0x73,
	0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x61,
	0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1a,
	0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xde, 0x01, 0x0a, 0x0a, 0x4b,
	0x69, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32,
	0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x76, 0x65,
	0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x23, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x12, 0x23, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74,
	0x49, 0x64, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x03, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x76,
	0x65, 0x6c, 0x5f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x76, 0x65, 0x6c, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xf9, 0x01, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22,
	0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x62, 0x61, 0x6c,
	0x6c, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x6f, 0x62, 0x6f,
	0x74, 0x52, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0b, 0x6b, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x0a, 0x6b, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2a, 0x92, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x50, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x22, 0x0a, 0x1e, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x4c, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x41, 0x4c, 0x4c,
	0x53, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c,
	0x45, 0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x50,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x5f, 0x4b,
	0x49, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x03, 0x42, 0x63, 0x42,
	0x1e, 0x53, 0x73, 0x6c, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f,
	0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72,
}

var (
	file_tracker_ssl_vision_detection_tracked_proto_rawDescOnce sync.Once
	file_tracker_ssl_vision_detection_tracked_proto_rawDescData = file_tracker_ssl_vision_detection_tracked_proto_rawDesc
)

func file_tracker_ssl_vision_detection_tracked_proto_rawDescGZIP() []byte {
	file_tracker_ssl_vision_detection_tracked_proto_rawDescOnce.Do(func() {
		file_tracker_ssl_vision_detection_tracked_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracker_ssl_vision_detection_tracked_proto_rawDescData)
	})
	return file_tracker_ssl_vision_detection_tracked_proto_rawDescData
}

var file_tracker_ssl_vision_detection_tracked_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tracker_ssl_vision_detection_tracked_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tracker_ssl_vision_detection_tracked_proto_goTypes = []any{
	(Capability)(0),       // 0: Capability
	(*TrackedBall)(nil),   // 1: TrackedBall
	(*KickedBall)(nil),    // 2: KickedBall
	(*TrackedRobot)(nil),  // 3: TrackedRobot
	(*TrackedFrame)(nil),  // 4: TrackedFrame
	(*geom.Vector3)(nil),  // 5: Vector3
	(*geom.Vector2)(nil),  // 6: Vector2
	(*state.RobotId)(nil), // 7: RobotId
}
var file_tracker_ssl_vision_detection_tracked_proto_depIdxs = []int32{
	5,  // 0: TrackedBall.pos:type_name -> Vector3
	5,  // 1: TrackedBall.vel:type_name -> Vector3
	6,  // 2: KickedBall.pos:type_name -> Vector2
	5,  // 3: KickedBall.vel:type_name -> Vector3
	6,  // 4: KickedBall.stop_pos:type_name -> Vector2
	7,  // 5: KickedBall.robot_id:type_name -> RobotId
	7,  // 6: TrackedRobot.robot_id:type_name -> RobotId
	6,  // 7: TrackedRobot.pos:type_name -> Vector2
	6,  // 8: TrackedRobot.vel:type_name -> Vector2
	1,  // 9: TrackedFrame.balls:type_name -> TrackedBall
	3,  // 10: TrackedFrame.robots:type_name -> TrackedRobot
	2,  // 11: TrackedFrame.kicked_ball:type_name -> KickedBall
	0,  // 12: TrackedFrame.capabilities:type_name -> Capability
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_tracker_ssl_vision_detection_tracked_proto_init() }
func file_tracker_ssl_vision_detection_tracked_proto_init() {
	if File_tracker_ssl_vision_detection_tracked_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tracker_ssl_vision_detection_tracked_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracker_ssl_vision_detection_tracked_proto_goTypes,
		DependencyIndexes: file_tracker_ssl_vision_detection_tracked_proto_depIdxs,
		EnumInfos:         file_tracker_ssl_vision_detection_tracked_proto_enumTypes,
		MessageInfos:      file_tracker_ssl_vision_detection_tracked_proto_msgTypes,
	}.Build()
	File_tracker_ssl_vision_detection_tracked_proto = out.File
	file_tracker_ssl_vision_detection_tracked_proto_rawDesc = nil
	file_tracker_ssl_vision_detection_tracked_proto_goTypes = nil
	file_tracker_ssl_vision_detection_tracked_proto_depIdxs = nil
}
