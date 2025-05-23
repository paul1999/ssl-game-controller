// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: rcon/ssl_gc_rcon_remotecontrol.proto

package rcon

import (
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

// All possible request types that the remote control can make
type RemoteControlRequestType int32

const (
	RemoteControlRequestType_UNKNOWN_REQUEST_TYPE RemoteControlRequestType = 0
	RemoteControlRequestType_EMERGENCY_STOP       RemoteControlRequestType = 1
	RemoteControlRequestType_ROBOT_SUBSTITUTION   RemoteControlRequestType = 2
	RemoteControlRequestType_TIMEOUT              RemoteControlRequestType = 3
	RemoteControlRequestType_CHALLENGE_FLAG       RemoteControlRequestType = 4
	RemoteControlRequestType_CHANGE_KEEPER_ID     RemoteControlRequestType = 5
	RemoteControlRequestType_STOP_TIMEOUT         RemoteControlRequestType = 6
	RemoteControlRequestType_FAIL_BALLPLACEMENT   RemoteControlRequestType = 7
)

// Enum value maps for RemoteControlRequestType.
var (
	RemoteControlRequestType_name = map[int32]string{
		0: "UNKNOWN_REQUEST_TYPE",
		1: "EMERGENCY_STOP",
		2: "ROBOT_SUBSTITUTION",
		3: "TIMEOUT",
		4: "CHALLENGE_FLAG",
		5: "CHANGE_KEEPER_ID",
		6: "STOP_TIMEOUT",
		7: "FAIL_BALLPLACEMENT",
	}
	RemoteControlRequestType_value = map[string]int32{
		"UNKNOWN_REQUEST_TYPE": 0,
		"EMERGENCY_STOP":       1,
		"ROBOT_SUBSTITUTION":   2,
		"TIMEOUT":              3,
		"CHALLENGE_FLAG":       4,
		"CHANGE_KEEPER_ID":     5,
		"STOP_TIMEOUT":         6,
		"FAIL_BALLPLACEMENT":   7,
	}
)

func (x RemoteControlRequestType) Enum() *RemoteControlRequestType {
	p := new(RemoteControlRequestType)
	*p = x
	return p
}

func (x RemoteControlRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteControlRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes[0].Descriptor()
}

func (RemoteControlRequestType) Type() protoreflect.EnumType {
	return &file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes[0]
}

func (x RemoteControlRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RemoteControlRequestType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RemoteControlRequestType(num)
	return nil
}

// Deprecated: Use RemoteControlRequestType.Descriptor instead.
func (RemoteControlRequestType) EnumDescriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{0}
}

type RemoteControlToController_Request int32

const (
	RemoteControlToController_UNKNOWN RemoteControlToController_Request = 0
	// Ping the GC to test the connection. The GC will respond with OK and the current team state
	RemoteControlToController_PING RemoteControlToController_Request = 1
	// Raise a challenge flag (this is not revocable)
	RemoteControlToController_CHALLENGE_FLAG RemoteControlToController_Request = 2
	// Stop an ongoing timeout
	RemoteControlToController_STOP_TIMEOUT RemoteControlToController_Request = 3
	// Fail a ball placement
	RemoteControlToController_FAIL_BALLPLACEMENT RemoteControlToController_Request = 4
)

// Enum value maps for RemoteControlToController_Request.
var (
	RemoteControlToController_Request_name = map[int32]string{
		0: "UNKNOWN",
		1: "PING",
		2: "CHALLENGE_FLAG",
		3: "STOP_TIMEOUT",
		4: "FAIL_BALLPLACEMENT",
	}
	RemoteControlToController_Request_value = map[string]int32{
		"UNKNOWN":            0,
		"PING":               1,
		"CHALLENGE_FLAG":     2,
		"STOP_TIMEOUT":       3,
		"FAIL_BALLPLACEMENT": 4,
	}
)

func (x RemoteControlToController_Request) Enum() *RemoteControlToController_Request {
	p := new(RemoteControlToController_Request)
	*p = x
	return p
}

func (x RemoteControlToController_Request) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteControlToController_Request) Descriptor() protoreflect.EnumDescriptor {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes[1].Descriptor()
}

func (RemoteControlToController_Request) Type() protoreflect.EnumType {
	return &file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes[1]
}

func (x RemoteControlToController_Request) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RemoteControlToController_Request) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RemoteControlToController_Request(num)
	return nil
}

// Deprecated: Use RemoteControlToController_Request.Descriptor instead.
func (RemoteControlToController_Request) EnumDescriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{1, 0}
}

// a registration that must be send by the remote control to the controller as the very first message
type RemoteControlRegistration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the team to be controlled
	Team *state.Team `protobuf:"varint,1,req,name=team,enum=Team" json:"team,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature     *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoteControlRegistration) Reset() {
	*x = RemoteControlRegistration{}
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoteControlRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteControlRegistration) ProtoMessage() {}

func (x *RemoteControlRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteControlRegistration.ProtoReflect.Descriptor instead.
func (*RemoteControlRegistration) Descriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteControlRegistration) GetTeam() state.Team {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return state.Team(0)
}

func (x *RemoteControlRegistration) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// wrapper for all messages from the remote control to the controller
type RemoteControlToController struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Types that are valid to be assigned to Msg:
	//
	//	*RemoteControlToController_Request_
	//	*RemoteControlToController_DesiredKeeper
	//	*RemoteControlToController_RequestRobotSubstitution
	//	*RemoteControlToController_RequestTimeout
	//	*RemoteControlToController_RequestEmergencyStop
	Msg           isRemoteControlToController_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoteControlToController) Reset() {
	*x = RemoteControlToController{}
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoteControlToController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteControlToController) ProtoMessage() {}

func (x *RemoteControlToController) ProtoReflect() protoreflect.Message {
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteControlToController.ProtoReflect.Descriptor instead.
func (*RemoteControlToController) Descriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteControlToController) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *RemoteControlToController) GetMsg() isRemoteControlToController_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *RemoteControlToController) GetRequest() RemoteControlToController_Request {
	if x != nil {
		if x, ok := x.Msg.(*RemoteControlToController_Request_); ok {
			return x.Request
		}
	}
	return RemoteControlToController_UNKNOWN
}

func (x *RemoteControlToController) GetDesiredKeeper() int32 {
	if x != nil {
		if x, ok := x.Msg.(*RemoteControlToController_DesiredKeeper); ok {
			return x.DesiredKeeper
		}
	}
	return 0
}

func (x *RemoteControlToController) GetRequestRobotSubstitution() bool {
	if x != nil {
		if x, ok := x.Msg.(*RemoteControlToController_RequestRobotSubstitution); ok {
			return x.RequestRobotSubstitution
		}
	}
	return false
}

func (x *RemoteControlToController) GetRequestTimeout() bool {
	if x != nil {
		if x, ok := x.Msg.(*RemoteControlToController_RequestTimeout); ok {
			return x.RequestTimeout
		}
	}
	return false
}

func (x *RemoteControlToController) GetRequestEmergencyStop() bool {
	if x != nil {
		if x, ok := x.Msg.(*RemoteControlToController_RequestEmergencyStop); ok {
			return x.RequestEmergencyStop
		}
	}
	return false
}

type isRemoteControlToController_Msg interface {
	isRemoteControlToController_Msg()
}

type RemoteControlToController_Request_ struct {
	// send a request to the GC
	Request RemoteControlToController_Request `protobuf:"varint,2,opt,name=request,enum=RemoteControlToController_Request,oneof"`
}

type RemoteControlToController_DesiredKeeper struct {
	// request a new desired keeper id
	DesiredKeeper int32 `protobuf:"varint,3,opt,name=desired_keeper,json=desiredKeeper,oneof"`
}

type RemoteControlToController_RequestRobotSubstitution struct {
	// true: request to substitute a robot at the next possibility
	// false: cancel request
	RequestRobotSubstitution bool `protobuf:"varint,4,opt,name=request_robot_substitution,json=requestRobotSubstitution,oneof"`
}

type RemoteControlToController_RequestTimeout struct {
	// true: request a timeout with the next stoppage
	// false: cancel the request
	RequestTimeout bool `protobuf:"varint,5,opt,name=request_timeout,json=requestTimeout,oneof"`
}

type RemoteControlToController_RequestEmergencyStop struct {
	// true: initiate an emergency stop
	// false: cancel the request
	RequestEmergencyStop bool `protobuf:"varint,6,opt,name=request_emergency_stop,json=requestEmergencyStop,oneof"`
}

func (*RemoteControlToController_Request_) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_DesiredKeeper) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_RequestRobotSubstitution) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_RequestTimeout) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_RequestEmergencyStop) isRemoteControlToController_Msg() {}

// wrapper for all messages from controller to a team's computer
type ControllerToRemoteControl struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply" json:"controller_reply,omitempty"`
	// current team state
	State         *RemoteControlTeamState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ControllerToRemoteControl) Reset() {
	*x = ControllerToRemoteControl{}
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControllerToRemoteControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerToRemoteControl) ProtoMessage() {}

func (x *ControllerToRemoteControl) ProtoReflect() protoreflect.Message {
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerToRemoteControl.ProtoReflect.Descriptor instead.
func (*ControllerToRemoteControl) Descriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{2}
}

func (x *ControllerToRemoteControl) GetControllerReply() *ControllerReply {
	if x != nil {
		return x.ControllerReply
	}
	return nil
}

func (x *ControllerToRemoteControl) GetState() *RemoteControlTeamState {
	if x != nil {
		return x.State
	}
	return nil
}

// Current team state from Controller for remote control
type RemoteControlTeamState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the team that is controlled
	Team *state.Team `protobuf:"varint,12,opt,name=team,enum=Team" json:"team,omitempty"`
	// list of all currently available request types that can be made
	AvailableRequests []RemoteControlRequestType `protobuf:"varint,1,rep,name=available_requests,json=availableRequests,enum=RemoteControlRequestType" json:"available_requests,omitempty"`
	// list of all currently active request types that are pending
	ActiveRequests []RemoteControlRequestType `protobuf:"varint,2,rep,name=active_requests,json=activeRequests,enum=RemoteControlRequestType" json:"active_requests,omitempty"`
	// currently set keeper id
	KeeperId *int32 `protobuf:"varint,3,opt,name=keeper_id,json=keeperId" json:"keeper_id,omitempty"`
	// number of seconds till emergency stop is executed
	// zero, if no emergency stop requested
	EmergencyStopIn *float32 `protobuf:"fixed32,4,opt,name=emergency_stop_in,json=emergencyStopIn" json:"emergency_stop_in,omitempty"`
	// number of timeouts left for the team
	TimeoutsLeft *int32 `protobuf:"varint,5,opt,name=timeouts_left,json=timeoutsLeft" json:"timeouts_left,omitempty"`
	// number of seconds left for timeout for the team
	TimeoutTimeLeft *float32 `protobuf:"fixed32,10,opt,name=timeout_time_left,json=timeoutTimeLeft" json:"timeout_time_left,omitempty"`
	// number of challenge flags left for the team
	ChallengeFlagsLeft *int32 `protobuf:"varint,6,opt,name=challenge_flags_left,json=challengeFlagsLeft" json:"challenge_flags_left,omitempty"`
	// max number of robots currently allowed
	MaxRobots *int32 `protobuf:"varint,7,opt,name=max_robots,json=maxRobots" json:"max_robots,omitempty"`
	// current number of robots visible on field
	RobotsOnField *int32 `protobuf:"varint,9,opt,name=robots_on_field,json=robotsOnField" json:"robots_on_field,omitempty"`
	// list of due times for each active yellow card (in seconds)
	YellowCardsDue []float32 `protobuf:"fixed32,8,rep,name=yellow_cards_due,json=yellowCardsDue" json:"yellow_cards_due,omitempty"`
	// if true, team is allowed to substitute robots
	CanSubstituteRobot *bool `protobuf:"varint,11,opt,name=can_substitute_robot,json=canSubstituteRobot" json:"can_substitute_robot,omitempty"`
	// number of bot substitutions left by the team in this halftime
	BotSubstitutionsLeft *uint32 `protobuf:"varint,13,opt,name=bot_substitutions_left,json=botSubstitutionsLeft" json:"bot_substitutions_left,omitempty"`
	// number of seconds left for current bot substitution
	BotSubstitutionTimeLeft *float32 `protobuf:"fixed32,14,opt,name=bot_substitution_time_left,json=botSubstitutionTimeLeft" json:"bot_substitution_time_left,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *RemoteControlTeamState) Reset() {
	*x = RemoteControlTeamState{}
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoteControlTeamState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteControlTeamState) ProtoMessage() {}

func (x *RemoteControlTeamState) ProtoReflect() protoreflect.Message {
	mi := &file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteControlTeamState.ProtoReflect.Descriptor instead.
func (*RemoteControlTeamState) Descriptor() ([]byte, []int) {
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteControlTeamState) GetTeam() state.Team {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return state.Team(0)
}

func (x *RemoteControlTeamState) GetAvailableRequests() []RemoteControlRequestType {
	if x != nil {
		return x.AvailableRequests
	}
	return nil
}

func (x *RemoteControlTeamState) GetActiveRequests() []RemoteControlRequestType {
	if x != nil {
		return x.ActiveRequests
	}
	return nil
}

func (x *RemoteControlTeamState) GetKeeperId() int32 {
	if x != nil && x.KeeperId != nil {
		return *x.KeeperId
	}
	return 0
}

func (x *RemoteControlTeamState) GetEmergencyStopIn() float32 {
	if x != nil && x.EmergencyStopIn != nil {
		return *x.EmergencyStopIn
	}
	return 0
}

func (x *RemoteControlTeamState) GetTimeoutsLeft() int32 {
	if x != nil && x.TimeoutsLeft != nil {
		return *x.TimeoutsLeft
	}
	return 0
}

func (x *RemoteControlTeamState) GetTimeoutTimeLeft() float32 {
	if x != nil && x.TimeoutTimeLeft != nil {
		return *x.TimeoutTimeLeft
	}
	return 0
}

func (x *RemoteControlTeamState) GetChallengeFlagsLeft() int32 {
	if x != nil && x.ChallengeFlagsLeft != nil {
		return *x.ChallengeFlagsLeft
	}
	return 0
}

func (x *RemoteControlTeamState) GetMaxRobots() int32 {
	if x != nil && x.MaxRobots != nil {
		return *x.MaxRobots
	}
	return 0
}

func (x *RemoteControlTeamState) GetRobotsOnField() int32 {
	if x != nil && x.RobotsOnField != nil {
		return *x.RobotsOnField
	}
	return 0
}

func (x *RemoteControlTeamState) GetYellowCardsDue() []float32 {
	if x != nil {
		return x.YellowCardsDue
	}
	return nil
}

func (x *RemoteControlTeamState) GetCanSubstituteRobot() bool {
	if x != nil && x.CanSubstituteRobot != nil {
		return *x.CanSubstituteRobot
	}
	return false
}

func (x *RemoteControlTeamState) GetBotSubstitutionsLeft() uint32 {
	if x != nil && x.BotSubstitutionsLeft != nil {
		return *x.BotSubstitutionsLeft
	}
	return 0
}

func (x *RemoteControlTeamState) GetBotSubstitutionTimeLeft() float32 {
	if x != nil && x.BotSubstitutionTimeLeft != nil {
		return *x.BotSubstitutionTimeLeft
	}
	return 0
}

var File_rcon_ssl_gc_rcon_remotecontrol_proto protoreflect.FileDescriptor

var file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x63, 0x6f, 0x6e, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x73,
	0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x72, 0x63, 0x6f, 0x6e, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72,
	0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x19, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb8, 0x03, 0x0a, 0x19,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x6f, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x1a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x62, 0x6f, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x74, 0x6f,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x22,
	0x5e, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x46,
	0x4c, 0x41, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x49, 0x4c, 0x5f,
	0x42, 0x41, 0x4c, 0x4c, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x87, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x3b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xa3, 0x05, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x48, 0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x42, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x65, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x4c, 0x65,
	0x66, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x30,
	0x0a, 0x14, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x4c, 0x65, 0x66, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73,
	0x4f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x79, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x75, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x0e, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73, 0x44, 0x75,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x63, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x14, 0x62, 0x6f, 0x74, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x62, 0x6f, 0x74,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x17, 0x62,
	0x6f, 0x74, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x2a, 0xc1, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x54,
	0x49, 0x54, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45,
	0x4e, 0x47, 0x45, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4b, 0x45, 0x45, 0x50, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x05,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x42, 0x5d, 0x42, 0x1b, 0x53, 0x73,
	0x6c, 0x47, 0x63, 0x52, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d,
	0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescOnce sync.Once
	file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescData = file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDesc
)

func file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP() []byte {
	file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescOnce.Do(func() {
		file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescData)
	})
	return file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDescData
}

var file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rcon_ssl_gc_rcon_remotecontrol_proto_goTypes = []any{
	(RemoteControlRequestType)(0),          // 0: RemoteControlRequestType
	(RemoteControlToController_Request)(0), // 1: RemoteControlToController.Request
	(*RemoteControlRegistration)(nil),      // 2: RemoteControlRegistration
	(*RemoteControlToController)(nil),      // 3: RemoteControlToController
	(*ControllerToRemoteControl)(nil),      // 4: ControllerToRemoteControl
	(*RemoteControlTeamState)(nil),         // 5: RemoteControlTeamState
	(state.Team)(0),                        // 6: Team
	(*Signature)(nil),                      // 7: Signature
	(*ControllerReply)(nil),                // 8: ControllerReply
}
var file_rcon_ssl_gc_rcon_remotecontrol_proto_depIdxs = []int32{
	6, // 0: RemoteControlRegistration.team:type_name -> Team
	7, // 1: RemoteControlRegistration.signature:type_name -> Signature
	7, // 2: RemoteControlToController.signature:type_name -> Signature
	1, // 3: RemoteControlToController.request:type_name -> RemoteControlToController.Request
	8, // 4: ControllerToRemoteControl.controller_reply:type_name -> ControllerReply
	5, // 5: ControllerToRemoteControl.state:type_name -> RemoteControlTeamState
	6, // 6: RemoteControlTeamState.team:type_name -> Team
	0, // 7: RemoteControlTeamState.available_requests:type_name -> RemoteControlRequestType
	0, // 8: RemoteControlTeamState.active_requests:type_name -> RemoteControlRequestType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_rcon_ssl_gc_rcon_remotecontrol_proto_init() }
func file_rcon_ssl_gc_rcon_remotecontrol_proto_init() {
	if File_rcon_ssl_gc_rcon_remotecontrol_proto != nil {
		return
	}
	file_rcon_ssl_gc_rcon_proto_init()
	file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes[1].OneofWrappers = []any{
		(*RemoteControlToController_Request_)(nil),
		(*RemoteControlToController_DesiredKeeper)(nil),
		(*RemoteControlToController_RequestRobotSubstitution)(nil),
		(*RemoteControlToController_RequestTimeout)(nil),
		(*RemoteControlToController_RequestEmergencyStop)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rcon_ssl_gc_rcon_remotecontrol_proto_goTypes,
		DependencyIndexes: file_rcon_ssl_gc_rcon_remotecontrol_proto_depIdxs,
		EnumInfos:         file_rcon_ssl_gc_rcon_remotecontrol_proto_enumTypes,
		MessageInfos:      file_rcon_ssl_gc_rcon_remotecontrol_proto_msgTypes,
	}.Build()
	File_rcon_ssl_gc_rcon_remotecontrol_proto = out.File
	file_rcon_ssl_gc_rcon_remotecontrol_proto_rawDesc = nil
	file_rcon_ssl_gc_rcon_remotecontrol_proto_goTypes = nil
	file_rcon_ssl_gc_rcon_remotecontrol_proto_depIdxs = nil
}
