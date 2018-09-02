// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_team.proto

package refproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TeamToControllerRequest_AdvantageResponse int32

const (
	// no choice -> will default to STOP
	TeamToControllerRequest_UNDECIDED TeamToControllerRequest_AdvantageResponse = 0
	// stop the game and handle the foul immediately
	TeamToControllerRequest_STOP TeamToControllerRequest_AdvantageResponse = 0
	// continue the game until the next stop of the game, then handle the foul
	TeamToControllerRequest_CONTINUE TeamToControllerRequest_AdvantageResponse = 1
)

var TeamToControllerRequest_AdvantageResponse_name = map[int32]string{
	0: "UNDECIDED",
	// Duplicate value: 0: "STOP",
	1: "CONTINUE",
}
var TeamToControllerRequest_AdvantageResponse_value = map[string]int32{
	"UNDECIDED": 0,
	"STOP":      0,
	"CONTINUE":  1,
}

func (x TeamToControllerRequest_AdvantageResponse) String() string {
	return proto.EnumName(TeamToControllerRequest_AdvantageResponse_name, int32(x))
}
func (TeamToControllerRequest_AdvantageResponse) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{1, 0}
}

type ControllerToTeamRequest_AdvantageChoice_Foul int32

const (
	// default value when not set
	ControllerToTeamRequest_AdvantageChoice_UNKNOWN ControllerToTeamRequest_AdvantageChoice_Foul = 0
	// an opponent has crashed with one of the teams bots
	ControllerToTeamRequest_AdvantageChoice_COLLISION ControllerToTeamRequest_AdvantageChoice_Foul = 1
)

var ControllerToTeamRequest_AdvantageChoice_Foul_name = map[int32]string{
	0: "UNKNOWN",
	1: "COLLISION",
}
var ControllerToTeamRequest_AdvantageChoice_Foul_value = map[string]int32{
	"UNKNOWN":   0,
	"COLLISION": 1,
}

func (x ControllerToTeamRequest_AdvantageChoice_Foul) String() string {
	return proto.EnumName(ControllerToTeamRequest_AdvantageChoice_Foul_name, int32(x))
}
func (ControllerToTeamRequest_AdvantageChoice_Foul) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{2, 0, 0}
}

// a registration that must be send by teams and autoRefs to the controller as the very first message
type TeamRegistration struct {
	// the exact team name as published by the game-controller
	TeamName             string   `protobuf:"bytes,1,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamRegistration) Reset()         { *m = TeamRegistration{} }
func (m *TeamRegistration) String() string { return proto.CompactTextString(m) }
func (*TeamRegistration) ProtoMessage()    {}
func (*TeamRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{0}
}
func (m *TeamRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamRegistration.Unmarshal(m, b)
}
func (m *TeamRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamRegistration.Marshal(b, m, deterministic)
}
func (dst *TeamRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamRegistration.Merge(dst, src)
}
func (m *TeamRegistration) XXX_Size() int {
	return xxx_messageInfo_TeamRegistration.Size(m)
}
func (m *TeamRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_TeamRegistration proto.InternalMessageInfo

func (m *TeamRegistration) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

// wrapper for all requests from a team's computer to the controller
type TeamToControllerRequest struct {
	// Types that are valid to be assigned to Request:
	//	*TeamToControllerRequest_DesiredKeeper
	//	*TeamToControllerRequest_AdvantageResponse_
	Request              isTeamToControllerRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TeamToControllerRequest) Reset()         { *m = TeamToControllerRequest{} }
func (m *TeamToControllerRequest) String() string { return proto.CompactTextString(m) }
func (*TeamToControllerRequest) ProtoMessage()    {}
func (*TeamToControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{1}
}
func (m *TeamToControllerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamToControllerRequest.Unmarshal(m, b)
}
func (m *TeamToControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamToControllerRequest.Marshal(b, m, deterministic)
}
func (dst *TeamToControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamToControllerRequest.Merge(dst, src)
}
func (m *TeamToControllerRequest) XXX_Size() int {
	return xxx_messageInfo_TeamToControllerRequest.Size(m)
}
func (m *TeamToControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamToControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamToControllerRequest proto.InternalMessageInfo

type isTeamToControllerRequest_Request interface {
	isTeamToControllerRequest_Request()
}

type TeamToControllerRequest_DesiredKeeper struct {
	DesiredKeeper *BotId `protobuf:"bytes,1,opt,name=desired_keeper,json=desiredKeeper,proto3,oneof"`
}

type TeamToControllerRequest_AdvantageResponse_ struct {
	AdvantageResponse TeamToControllerRequest_AdvantageResponse `protobuf:"varint,2,opt,name=advantage_response,json=advantageResponse,proto3,enum=TeamToControllerRequest_AdvantageResponse,oneof"`
}

func (*TeamToControllerRequest_DesiredKeeper) isTeamToControllerRequest_Request() {}

func (*TeamToControllerRequest_AdvantageResponse_) isTeamToControllerRequest_Request() {}

func (m *TeamToControllerRequest) GetRequest() isTeamToControllerRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *TeamToControllerRequest) GetDesiredKeeper() *BotId {
	if x, ok := m.GetRequest().(*TeamToControllerRequest_DesiredKeeper); ok {
		return x.DesiredKeeper
	}
	return nil
}

func (m *TeamToControllerRequest) GetAdvantageResponse() TeamToControllerRequest_AdvantageResponse {
	if x, ok := m.GetRequest().(*TeamToControllerRequest_AdvantageResponse_); ok {
		return x.AdvantageResponse
	}
	return TeamToControllerRequest_UNDECIDED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TeamToControllerRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TeamToControllerRequest_OneofMarshaler, _TeamToControllerRequest_OneofUnmarshaler, _TeamToControllerRequest_OneofSizer, []interface{}{
		(*TeamToControllerRequest_DesiredKeeper)(nil),
		(*TeamToControllerRequest_AdvantageResponse_)(nil),
	}
}

func _TeamToControllerRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TeamToControllerRequest)
	// request
	switch x := m.Request.(type) {
	case *TeamToControllerRequest_DesiredKeeper:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DesiredKeeper); err != nil {
			return err
		}
	case *TeamToControllerRequest_AdvantageResponse_:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.AdvantageResponse))
	case nil:
	default:
		return fmt.Errorf("TeamToControllerRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _TeamToControllerRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TeamToControllerRequest)
	switch tag {
	case 1: // request.desired_keeper
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BotId)
		err := b.DecodeMessage(msg)
		m.Request = &TeamToControllerRequest_DesiredKeeper{msg}
		return true, err
	case 2: // request.advantage_response
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Request = &TeamToControllerRequest_AdvantageResponse_{TeamToControllerRequest_AdvantageResponse(x)}
		return true, err
	default:
		return false, nil
	}
}

func _TeamToControllerRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TeamToControllerRequest)
	// request
	switch x := m.Request.(type) {
	case *TeamToControllerRequest_DesiredKeeper:
		s := proto.Size(x.DesiredKeeper)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TeamToControllerRequest_AdvantageResponse_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.AdvantageResponse))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// wrapper for all requests from controller to a team's computer
type ControllerToTeamRequest struct {
	// Types that are valid to be assigned to Request:
	//	*ControllerToTeamRequest_AdvantageChoice_
	Request              isControllerToTeamRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ControllerToTeamRequest) Reset()         { *m = ControllerToTeamRequest{} }
func (m *ControllerToTeamRequest) String() string { return proto.CompactTextString(m) }
func (*ControllerToTeamRequest) ProtoMessage()    {}
func (*ControllerToTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{2}
}
func (m *ControllerToTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerToTeamRequest.Unmarshal(m, b)
}
func (m *ControllerToTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerToTeamRequest.Marshal(b, m, deterministic)
}
func (dst *ControllerToTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerToTeamRequest.Merge(dst, src)
}
func (m *ControllerToTeamRequest) XXX_Size() int {
	return xxx_messageInfo_ControllerToTeamRequest.Size(m)
}
func (m *ControllerToTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerToTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerToTeamRequest proto.InternalMessageInfo

type isControllerToTeamRequest_Request interface {
	isControllerToTeamRequest_Request()
}

type ControllerToTeamRequest_AdvantageChoice_ struct {
	AdvantageChoice *ControllerToTeamRequest_AdvantageChoice `protobuf:"bytes,1,opt,name=advantage_choice,json=advantageChoice,proto3,oneof"`
}

func (*ControllerToTeamRequest_AdvantageChoice_) isControllerToTeamRequest_Request() {}

func (m *ControllerToTeamRequest) GetRequest() isControllerToTeamRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ControllerToTeamRequest) GetAdvantageChoice() *ControllerToTeamRequest_AdvantageChoice {
	if x, ok := m.GetRequest().(*ControllerToTeamRequest_AdvantageChoice_); ok {
		return x.AdvantageChoice
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ControllerToTeamRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ControllerToTeamRequest_OneofMarshaler, _ControllerToTeamRequest_OneofUnmarshaler, _ControllerToTeamRequest_OneofSizer, []interface{}{
		(*ControllerToTeamRequest_AdvantageChoice_)(nil),
	}
}

func _ControllerToTeamRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ControllerToTeamRequest)
	// request
	switch x := m.Request.(type) {
	case *ControllerToTeamRequest_AdvantageChoice_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AdvantageChoice); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ControllerToTeamRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _ControllerToTeamRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ControllerToTeamRequest)
	switch tag {
	case 1: // request.advantage_choice
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControllerToTeamRequest_AdvantageChoice)
		err := b.DecodeMessage(msg)
		m.Request = &ControllerToTeamRequest_AdvantageChoice_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ControllerToTeamRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ControllerToTeamRequest)
	// request
	switch x := m.Request.(type) {
	case *ControllerToTeamRequest_AdvantageChoice_:
		s := proto.Size(x.AdvantageChoice)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ControllerToTeamRequest_AdvantageChoice struct {
	// the type of foul that occurred
	Foul                 ControllerToTeamRequest_AdvantageChoice_Foul `protobuf:"varint,1,opt,name=foul,proto3,enum=ControllerToTeamRequest_AdvantageChoice_Foul" json:"foul,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *ControllerToTeamRequest_AdvantageChoice) Reset() {
	*m = ControllerToTeamRequest_AdvantageChoice{}
}
func (m *ControllerToTeamRequest_AdvantageChoice) String() string { return proto.CompactTextString(m) }
func (*ControllerToTeamRequest_AdvantageChoice) ProtoMessage()    {}
func (*ControllerToTeamRequest_AdvantageChoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_game_controller_team_39dad039598008d6, []int{2, 0}
}
func (m *ControllerToTeamRequest_AdvantageChoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice.Unmarshal(m, b)
}
func (m *ControllerToTeamRequest_AdvantageChoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice.Marshal(b, m, deterministic)
}
func (dst *ControllerToTeamRequest_AdvantageChoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice.Merge(dst, src)
}
func (m *ControllerToTeamRequest_AdvantageChoice) XXX_Size() int {
	return xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice.Size(m)
}
func (m *ControllerToTeamRequest_AdvantageChoice) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerToTeamRequest_AdvantageChoice proto.InternalMessageInfo

func (m *ControllerToTeamRequest_AdvantageChoice) GetFoul() ControllerToTeamRequest_AdvantageChoice_Foul {
	if m != nil {
		return m.Foul
	}
	return ControllerToTeamRequest_AdvantageChoice_UNKNOWN
}

func init() {
	proto.RegisterType((*TeamRegistration)(nil), "TeamRegistration")
	proto.RegisterType((*TeamToControllerRequest)(nil), "TeamToControllerRequest")
	proto.RegisterType((*ControllerToTeamRequest)(nil), "ControllerToTeamRequest")
	proto.RegisterType((*ControllerToTeamRequest_AdvantageChoice)(nil), "ControllerToTeamRequest.AdvantageChoice")
	proto.RegisterEnum("TeamToControllerRequest_AdvantageResponse", TeamToControllerRequest_AdvantageResponse_name, TeamToControllerRequest_AdvantageResponse_value)
	proto.RegisterEnum("ControllerToTeamRequest_AdvantageChoice_Foul", ControllerToTeamRequest_AdvantageChoice_Foul_name, ControllerToTeamRequest_AdvantageChoice_Foul_value)
}

func init() {
	proto.RegisterFile("ssl_game_controller_team.proto", fileDescriptor_ssl_game_controller_team_39dad039598008d6)
}

var fileDescriptor_ssl_game_controller_team_39dad039598008d6 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0xcb, 0xda, 0x40,
	0x10, 0xc5, 0x13, 0x11, 0x35, 0x63, 0xd5, 0xb8, 0x17, 0xc5, 0x42, 0x91, 0x9c, 0xa4, 0xd0, 0x08,
	0xf6, 0x5e, 0xd0, 0x68, 0x31, 0x28, 0x9b, 0xb2, 0x26, 0xf4, 0xd0, 0x43, 0xd8, 0x26, 0x53, 0x1b,
	0x9a, 0x64, 0xed, 0x66, 0x2d, 0xfd, 0xbb, 0x0b, 0xbd, 0x97, 0x24, 0xb6, 0xa2, 0xfd, 0x84, 0xef,
	0x98, 0x79, 0xef, 0xf7, 0x26, 0x6f, 0x58, 0x78, 0x55, 0x14, 0x69, 0x78, 0xe4, 0x19, 0x86, 0x91,
	0xc8, 0x95, 0x14, 0x69, 0x8a, 0x32, 0x54, 0xc8, 0x33, 0xfb, 0x24, 0x85, 0x12, 0x93, 0xe9, 0x53,
	0x7a, 0x24, 0xb2, 0x4c, 0xe4, 0xb5, 0xc3, 0x9a, 0x83, 0xe9, 0x23, 0xcf, 0x18, 0x1e, 0x93, 0x42,
	0x49, 0xae, 0x12, 0x91, 0x93, 0x97, 0x60, 0x94, 0x19, 0x61, 0xce, 0x33, 0x1c, 0xeb, 0x53, 0x7d,
	0x66, 0xb0, 0x4e, 0x39, 0xa0, 0x3c, 0x43, 0xeb, 0x97, 0x0e, 0xa3, 0x92, 0xf0, 0x85, 0xf3, 0x2f,
	0x92, 0xe1, 0xf7, 0x33, 0x16, 0x8a, 0xcc, 0xa1, 0x1f, 0x63, 0x91, 0x48, 0x8c, 0xc3, 0x6f, 0x88,
	0x27, 0x94, 0x15, 0xdd, 0x5d, 0xb4, 0xec, 0x95, 0x50, 0x6e, 0xbc, 0xd5, 0x58, 0xef, 0xa2, 0xef,
	0x2a, 0x99, 0x7c, 0x02, 0xc2, 0xe3, 0x1f, 0x3c, 0x57, 0xfc, 0x88, 0xa1, 0xc4, 0xe2, 0x24, 0xf2,
	0x02, 0xc7, 0x8d, 0xa9, 0x3e, 0xeb, 0x2f, 0x5e, 0xdb, 0x0f, 0xd6, 0xd8, 0xcb, 0xbf, 0x08, 0xbb,
	0x10, 0x5b, 0x8d, 0x0d, 0xf9, 0xfd, 0xd0, 0x7a, 0x07, 0xc3, 0xff, 0x9c, 0xa4, 0x07, 0x46, 0x40,
	0xd7, 0x1b, 0xc7, 0x5d, 0x6f, 0xd6, 0xa6, 0x46, 0x3a, 0xd0, 0x3c, 0xf8, 0xde, 0x07, 0x53, 0x23,
	0x2f, 0xa0, 0xe3, 0x78, 0xd4, 0x77, 0x69, 0xb0, 0x31, 0xf5, 0x49, 0xc3, 0xd4, 0x57, 0x06, 0xb4,
	0x65, 0xbd, 0xd1, 0xfa, 0xad, 0xc3, 0xe8, 0xfa, 0x1f, 0xbe, 0xa8, 0x4f, 0x56, 0x97, 0x0e, 0xc0,
	0xbc, 0x76, 0x88, 0xbe, 0x8a, 0x24, 0xc2, 0x4b, 0xed, 0x99, 0xfd, 0x80, 0xb9, 0x36, 0x70, 0x2a,
	0xff, 0x56, 0x63, 0x03, 0x7e, 0x3b, 0x9a, 0xfc, 0x84, 0xc1, 0x9d, 0x8b, 0x2c, 0xa1, 0xf9, 0x45,
	0x9c, 0xd3, 0x2a, 0xbd, 0xbf, 0x78, 0xf3, 0xdc, 0x74, 0xfb, 0xbd, 0x38, 0xa7, 0xac, 0x42, 0x2d,
	0x0b, 0x9a, 0xe5, 0x17, 0xe9, 0x42, 0x3b, 0xa0, 0x3b, 0xea, 0x7d, 0xa4, 0xa6, 0x56, 0xde, 0xc4,
	0xf1, 0xf6, 0x7b, 0xf7, 0xe0, 0x7a, 0xf4, 0xa6, 0xf7, 0xe7, 0x56, 0xf5, 0x48, 0xde, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0xe9, 0x88, 0x26, 0x68, 0x02, 0x00, 0x00,
}