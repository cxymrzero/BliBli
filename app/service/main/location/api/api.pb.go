// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/location/api/v1/api.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		app/service/main/location/api/v1/api.proto

	It has these top-level messages:
		InfoReply
		InfoReq
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import binary "encoding/binary"

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

type InfoReply struct {
	Addr      string  `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Country   string  `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Province  string  `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City      string  `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Isp       string  `protobuf:"bytes,5,opt,name=isp,proto3" json:"isp,omitempty"`
	Latitude  float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	ZoneId    int64   `protobuf:"varint,8,opt,name=zone_id,json=zoneId,proto3" json:"zoneId,omitempty"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

func (m *InfoReply) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *InfoReply) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *InfoReply) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *InfoReply) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *InfoReply) GetIsp() string {
	if m != nil {
		return m.Isp
	}
	return ""
}

func (m *InfoReply) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *InfoReply) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *InfoReply) GetZoneId() int64 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

type InfoReq struct {
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *InfoReq) Reset()                    { *m = InfoReq{} }
func (m *InfoReq) String() string            { return proto.CompactTextString(m) }
func (*InfoReq) ProtoMessage()               {}
func (*InfoReq) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *InfoReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoReply)(nil), "location.service.InfoReply")
	proto.RegisterType((*InfoReq)(nil), "location.service.InfoReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Location service

type LocationClient interface {
	// Info get ip info.
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoReply, error)
}

type locationClient struct {
	cc *grpc.ClientConn
}

func NewLocationClient(cc *grpc.ClientConn) LocationClient {
	return &locationClient{cc}
}

func (c *locationClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/location.service.Location/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Location service

type LocationServer interface {
	// Info get ip info.
	Info(context.Context, *InfoReq) (*InfoReply, error)
}

func RegisterLocationServer(s *grpc.Server, srv LocationServer) {
	s.RegisterService(&_Location_serviceDesc, srv)
}

func _Location_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/location.service.Location/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Location_serviceDesc = grpc.ServiceDesc{
	ServiceName: "location.service.Location",
	HandlerType: (*LocationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Location_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/main/location/api/v1/api.proto",
}

func (m *InfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if len(m.Country) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Country)))
		i += copy(dAtA[i:], m.Country)
	}
	if len(m.Province) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Province)))
		i += copy(dAtA[i:], m.Province)
	}
	if len(m.City) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.City)))
		i += copy(dAtA[i:], m.City)
	}
	if len(m.Isp) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Isp)))
		i += copy(dAtA[i:], m.Isp)
	}
	if m.Latitude != 0 {
		dAtA[i] = 0x31
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i += 8
	}
	if m.Longitude != 0 {
		dAtA[i] = 0x39
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i += 8
	}
	if m.ZoneId != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.ZoneId))
	}
	return i, nil
}

func (m *InfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InfoReply) Size() (n int) {
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Province)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Isp)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	if m.ZoneId != 0 {
		n += 1 + sovApi(uint64(m.ZoneId))
	}
	return n
}

func (m *InfoReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
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
func (m *InfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
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
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Province", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.Province = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Isp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.Isp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneId", wireType)
			}
			m.ZoneId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZoneId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
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
func (m *InfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
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
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
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
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
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
			b := dAtA[iNdEx]
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
				if dAtA[iNdEx-1] < 0x80 {
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
				b := dAtA[iNdEx]
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
					b := dAtA[iNdEx]
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
				next, err := skipApi(dAtA[start:])
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

func init() { proto.RegisterFile("app/service/main/location/api/v1/api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0x4d, 0x4a, 0xc4, 0x30,
	0x18, 0x35, 0xd3, 0xda, 0x9f, 0x6f, 0x35, 0x04, 0xc1, 0x38, 0x6a, 0x29, 0xb3, 0x2a, 0x2e, 0x5a,
	0x46, 0xf7, 0x2e, 0xdc, 0x8d, 0xb8, 0xea, 0xd2, 0x8d, 0xc4, 0x36, 0x4a, 0xa0, 0x26, 0x31, 0x93,
	0x29, 0xd4, 0x93, 0x78, 0x24, 0x57, 0xe2, 0x11, 0xa4, 0x5e, 0x44, 0x92, 0x76, 0x3a, 0x20, 0xb8,
	0xca, 0xfb, 0xc9, 0xf7, 0xe0, 0x3d, 0xb8, 0xa0, 0x4a, 0x15, 0x1b, 0xa6, 0x5b, 0x5e, 0xb1, 0xe2,
	0x85, 0x72, 0x51, 0x34, 0xb2, 0xa2, 0x86, 0x4b, 0x51, 0x50, 0xc5, 0x8b, 0x76, 0x65, 0x9f, 0x5c,
	0x69, 0x69, 0x24, 0x9e, 0xef, 0xac, 0x7c, 0x3c, 0x58, 0x7e, 0x22, 0x88, 0xd7, 0xe2, 0x49, 0x96,
	0x4c, 0x35, 0x1d, 0xc6, 0xe0, 0xd3, 0xba, 0xd6, 0x04, 0xa5, 0x28, 0x8b, 0x4b, 0x87, 0x31, 0x81,
	0xb0, 0x92, 0x5b, 0x61, 0x74, 0x47, 0x66, 0x4e, 0xde, 0x51, 0xbc, 0x80, 0x48, 0x69, 0xd9, 0x72,
	0x51, 0x31, 0xe2, 0x39, 0x6b, 0xe2, 0x36, 0xa9, 0xe2, 0xa6, 0x23, 0xfe, 0x90, 0x64, 0x31, 0x9e,
	0x83, 0xc7, 0x37, 0x8a, 0x1c, 0x3a, 0xc9, 0x42, 0x9b, 0xd0, 0x50, 0xc3, 0xcd, 0xb6, 0x66, 0x24,
	0x48, 0x51, 0x86, 0xca, 0x89, 0xe3, 0x33, 0x88, 0x1b, 0x29, 0x9e, 0x07, 0x33, 0x74, 0xe6, 0x5e,
	0xc0, 0xc7, 0x10, 0xbe, 0x49, 0xc1, 0x1e, 0x78, 0x4d, 0xa2, 0x14, 0x65, 0x5e, 0x19, 0x58, 0xba,
	0xae, 0x97, 0xe7, 0x10, 0x0e, 0x7d, 0x5e, 0xa7, 0x36, 0xb3, 0x7d, 0x9b, 0xcb, 0x5b, 0x88, 0xee,
	0xc6, 0x0d, 0xf0, 0x35, 0xf8, 0xf6, 0x2b, 0x3e, 0xc9, 0xff, 0xce, 0x92, 0x8f, 0x11, 0x8b, 0xd3,
	0xff, 0x2c, 0xd5, 0x74, 0x37, 0x47, 0x1f, 0x7d, 0x82, 0xbe, 0xfa, 0x04, 0x7d, 0xf7, 0x09, 0x7a,
	0xff, 0x49, 0x0e, 0xee, 0x67, 0xed, 0xea, 0x31, 0x70, 0x53, 0x5f, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x51, 0x4f, 0x85, 0x70, 0x98, 0x01, 0x00, 0x00,
}
