// Code generated by protoc-gen-gogo.
// source: api.proto
// DO NOT EDIT!

/*
	Package reporting is a generated protocol buffer package.

	It is generated from these files:
		api.proto

	It has these top-level messages:
		RawEvent
		RawEvents
*/
package reporting

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RawEvent a single event that happened in the system
type RawEvent struct {
	// Type is the event type such as "node accessed" or "user logged in"
	Type string `protobuf:"bytes,1,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	// Data is the JSON-encoded event payload
	Data []byte `protobuf:"bytes,2,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
	// Timestamp is the Unix timestamp of the event
	Timestamp int64 `protobuf:"varint,3,opt,name=Timestamp,json=timestamp,proto3" json:"Timestamp,omitempty"`
}

func (m *RawEvent) Reset()                    { *m = RawEvent{} }
func (m *RawEvent) String() string            { return proto.CompactTextString(m) }
func (*RawEvent) ProtoMessage()               {}
func (*RawEvent) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

// RawEvents defines a series of events
type RawEvents struct {
	// Events is a list of events
	Events []*RawEvent `protobuf:"bytes,1,rep,name=Events,json=events" json:"Events,omitempty"`
}

func (m *RawEvents) Reset()                    { *m = RawEvents{} }
func (m *RawEvents) String() string            { return proto.CompactTextString(m) }
func (*RawEvents) ProtoMessage()               {}
func (*RawEvents) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *RawEvents) GetEvents() []*RawEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*RawEvent)(nil), "reporting.RawEvent")
	proto.RegisterType((*RawEvents)(nil), "reporting.RawEvents")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Events service

type EventsClient interface {
	// Record records the provided list of events
	Record(ctx context.Context, in *RawEvents, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Record(ctx context.Context, in *RawEvents, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/reporting.Events/Record", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Events service

type EventsServer interface {
	// Record records the provided list of events
	Record(context.Context, *RawEvents) (*google_protobuf1.Empty, error)
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawEvents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reporting.Events/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Record(ctx, req.(*RawEvents))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reporting.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _Events_Record_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorApi,
}

func (m *RawEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RawEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintApi(data, i, uint64(len(m.Type)))
		i += copy(data[i:], m.Type)
	}
	if len(m.Data) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintApi(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.Timestamp != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintApi(data, i, uint64(m.Timestamp))
	}
	return i, nil
}

func (m *RawEvents) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RawEvents) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			data[i] = 0xa
			i++
			i = encodeVarintApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Api(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Api(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintApi(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RawEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovApi(uint64(m.Timestamp))
	}
	return n
}

func (m *RawEvents) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RawEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], data[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RawEvents) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawEvents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawEvents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &RawEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xdd, 0xb6, 0x04, 0xb3, 0x0a, 0xca, 0xfa, 0x41, 0xa8, 0x12, 0x42, 0x4e, 0x41, 0x71,
	0x17, 0xeb, 0x45, 0x3c, 0x8a, 0x3d, 0x0a, 0x12, 0x0a, 0x82, 0x17, 0x99, 0xa6, 0x6b, 0xba, 0x60,
	0xb2, 0xcb, 0xee, 0xb4, 0x25, 0x57, 0x5f, 0xc1, 0x8b, 0x8f, 0xe4, 0x51, 0xf0, 0x05, 0xa4, 0xfa,
	0x20, 0xd2, 0x26, 0xf1, 0xe4, 0x6d, 0x66, 0xfe, 0x1f, 0xc3, 0x8f, 0xfa, 0x60, 0x14, 0x37, 0x56,
	0xa3, 0x66, 0xbe, 0x95, 0x46, 0x5b, 0x54, 0x65, 0xde, 0x7f, 0xc8, 0x15, 0x4e, 0x67, 0x63, 0x9e,
	0xe9, 0x42, 0xe4, 0xd6, 0x64, 0x67, 0x32, 0xd3, 0xae, 0x72, 0x28, 0x9b, 0x35, 0x07, 0x94, 0x0b,
	0xa8, 0x04, 0x4e, 0x95, 0x9d, 0x3c, 0x1a, 0xb0, 0x58, 0x89, 0x5c, 0xeb, 0xfc, 0x59, 0x82, 0x51,
	0xae, 0x19, 0x05, 0x18, 0x25, 0xa0, 0x2c, 0x35, 0x02, 0x2a, 0x5d, 0xba, 0xfa, 0x4d, 0xff, 0xa8,
	0x51, 0xd7, 0xdb, 0x78, 0xf6, 0x24, 0x64, 0x61, 0xb0, 0xaa, 0xc5, 0xf8, 0x8e, 0x6e, 0xa6, 0xb0,
	0x18, 0xce, 0x65, 0x89, 0x8c, 0xd1, 0xde, 0xa8, 0x32, 0x32, 0x20, 0x11, 0x49, 0xfc, 0xb4, 0x87,
	0x95, 0x91, 0xab, 0xdb, 0x0d, 0x20, 0x04, 0x9d, 0x88, 0x24, 0xdb, 0x69, 0x6f, 0x02, 0x08, 0xec,
	0x98, 0xfa, 0x23, 0x55, 0x48, 0x87, 0x50, 0x98, 0xa0, 0x1b, 0x91, 0xa4, 0x9b, 0xfa, 0xd8, 0x1e,
	0xe2, 0x4b, 0xea, 0xb7, 0x8d, 0x8e, 0x9d, 0x52, 0xaf, 0x9e, 0x02, 0x12, 0x75, 0x93, 0xad, 0xc1,
	0x1e, 0xff, 0x63, 0xe6, 0xad, 0x2b, 0xf5, 0xe4, 0xda, 0x32, 0xb8, 0x6f, 0xcd, 0xec, 0x96, 0x7a,
	0xa9, 0xcc, 0xb4, 0x9d, 0xb0, 0xfd, 0x7f, 0x02, 0xae, 0x7f, 0xc8, 0x6b, 0x26, 0xde, 0x32, 0xf1,
	0xe1, 0x8a, 0x29, 0x3e, 0x78, 0xf9, 0xfc, 0x79, 0xed, 0xec, 0xc4, 0x54, 0xcc, 0xcf, 0x45, 0xdd,
	0x7a, 0x45, 0x4e, 0xae, 0x77, 0xdf, 0x97, 0x21, 0xf9, 0x58, 0x86, 0xe4, 0x6b, 0x19, 0x92, 0xb7,
	0xef, 0x70, 0x63, 0xec, 0xad, 0x83, 0x17, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0xb1, 0x60,
	0x4d, 0x8e, 0x01, 0x00, 0x00,
}
