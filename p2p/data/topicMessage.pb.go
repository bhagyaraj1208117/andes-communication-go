// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topicMessage.proto

package data

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TopicMessage struct {
	Version        uint32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Payload        []byte `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Timestamp      int64  `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Pk             []byte `protobuf:"bytes,4,opt,name=Pk,proto3" json:"Pk,omitempty"`
	SignatureOnPid []byte `protobuf:"bytes,5,opt,name=SignatureOnPid,proto3" json:"SignatureOnPid,omitempty"`
}

func (m *TopicMessage) Reset()      { *m = TopicMessage{} }
func (*TopicMessage) ProtoMessage() {}
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_131cdede10b420b6, []int{0}
}
func (m *TopicMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TopicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessage.Merge(m, src)
}
func (m *TopicMessage) XXX_Size() int {
	return m.Size()
}
func (m *TopicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessage proto.InternalMessageInfo

func (m *TopicMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TopicMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TopicMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TopicMessage) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *TopicMessage) GetSignatureOnPid() []byte {
	if m != nil {
		return m.SignatureOnPid
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicMessage)(nil), "proto.TopicMessage")
}

func init() { proto.RegisterFile("topicMessage.proto", fileDescriptor_131cdede10b420b6) }

var fileDescriptor_131cdede10b420b6 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xc9, 0x2f, 0xc8,
	0x4c, 0xf6, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9,
	0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba,
	0x94, 0x66, 0x30, 0x72, 0xf1, 0x84, 0x20, 0x19, 0x26, 0x24, 0xc1, 0xc5, 0x1e, 0x96, 0x5a, 0x54,
	0x9c, 0x99, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0x04, 0xe3, 0x82, 0x64, 0x02, 0x12,
	0x2b, 0x73, 0xf2, 0x13, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x21, 0x19,
	0x2e, 0xce, 0x90, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0xe6, 0x20, 0x84, 0x80, 0x10, 0x1f, 0x17, 0x53, 0x40, 0xb6, 0x04, 0x0b, 0x58, 0x0b, 0x53, 0x40,
	0xb6, 0x90, 0x1a, 0x17, 0x5f, 0x70, 0x66, 0x7a, 0x5e, 0x62, 0x49, 0x69, 0x51, 0xaa, 0x7f, 0x5e,
	0x40, 0x66, 0x8a, 0x04, 0x2b, 0x58, 0x0e, 0x4d, 0xd4, 0xc9, 0xee, 0xc2, 0x43, 0x39, 0x86, 0x1b,
	0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x49, 0x6c, 0x60, 0x1f,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x51, 0x72, 0x2b, 0x2d, 0x01, 0x00, 0x00,
}

func (this *TopicMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TopicMessage)
	if !ok {
		that2, ok := that.(TopicMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if !bytes.Equal(this.Pk, that1.Pk) {
		return false
	}
	if !bytes.Equal(this.SignatureOnPid, that1.SignatureOnPid) {
		return false
	}
	return true
}
func (this *TopicMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&data.TopicMessage{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Pk: "+fmt.Sprintf("%#v", this.Pk)+",\n")
	s = append(s, "SignatureOnPid: "+fmt.Sprintf("%#v", this.SignatureOnPid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTopicMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TopicMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopicMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignatureOnPid) > 0 {
		i -= len(m.SignatureOnPid)
		copy(dAtA[i:], m.SignatureOnPid)
		i = encodeVarintTopicMessage(dAtA, i, uint64(len(m.SignatureOnPid)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Pk) > 0 {
		i -= len(m.Pk)
		copy(dAtA[i:], m.Pk)
		i = encodeVarintTopicMessage(dAtA, i, uint64(len(m.Pk)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintTopicMessage(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTopicMessage(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintTopicMessage(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTopicMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovTopicMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TopicMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovTopicMessage(uint64(m.Version))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTopicMessage(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTopicMessage(uint64(m.Timestamp))
	}
	l = len(m.Pk)
	if l > 0 {
		n += 1 + l + sovTopicMessage(uint64(l))
	}
	l = len(m.SignatureOnPid)
	if l > 0 {
		n += 1 + l + sovTopicMessage(uint64(l))
	}
	return n
}

func sovTopicMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTopicMessage(x uint64) (n int) {
	return sovTopicMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TopicMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TopicMessage{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`Pk:` + fmt.Sprintf("%v", this.Pk) + `,`,
		`SignatureOnPid:` + fmt.Sprintf("%v", this.SignatureOnPid) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTopicMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TopicMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopicMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TopicMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTopicMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTopicMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTopicMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTopicMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pk = append(m.Pk[:0], dAtA[iNdEx:postIndex]...)
			if m.Pk == nil {
				m.Pk = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureOnPid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTopicMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTopicMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignatureOnPid = append(m.SignatureOnPid[:0], dAtA[iNdEx:postIndex]...)
			if m.SignatureOnPid == nil {
				m.SignatureOnPid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopicMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopicMessage
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
func skipTopicMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopicMessage
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
					return 0, ErrIntOverflowTopicMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTopicMessage
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
			if length < 0 {
				return 0, ErrInvalidLengthTopicMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTopicMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTopicMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTopicMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopicMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTopicMessage = fmt.Errorf("proto: unexpected end of group")
)
