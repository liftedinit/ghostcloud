// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ghostcloud/ghostcloud/payload.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type Payload struct {
	// Types that are valid to be assigned to PayloadOption:
	//	*Payload_Dataset
	//	*Payload_Archive
	PayloadOption isPayload_PayloadOption `protobuf_oneof:"payload_option"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a89c1701d582a59, []int{0}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

type isPayload_PayloadOption interface {
	isPayload_PayloadOption()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Payload_Dataset struct {
	Dataset *Dataset `protobuf:"bytes,1,opt,name=dataset,proto3,oneof" json:"dataset,omitempty"`
}
type Payload_Archive struct {
	Archive *Archive `protobuf:"bytes,2,opt,name=archive,proto3,oneof" json:"archive,omitempty"`
}

func (*Payload_Dataset) isPayload_PayloadOption() {}
func (*Payload_Archive) isPayload_PayloadOption() {}

func (m *Payload) GetPayloadOption() isPayload_PayloadOption {
	if m != nil {
		return m.PayloadOption
	}
	return nil
}

func (m *Payload) GetDataset() *Dataset {
	if x, ok := m.GetPayloadOption().(*Payload_Dataset); ok {
		return x.Dataset
	}
	return nil
}

func (m *Payload) GetArchive() *Archive {
	if x, ok := m.GetPayloadOption().(*Payload_Archive); ok {
		return x.Archive
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Payload_Dataset)(nil),
		(*Payload_Archive)(nil),
	}
}

func init() {
	proto.RegisterType((*Payload)(nil), "ghostcloud.ghostcloud.Payload")
}

func init() {
	proto.RegisterFile("ghostcloud/ghostcloud/payload.proto", fileDescriptor_0a89c1701d582a59)
}

var fileDescriptor_0a89c1701d582a59 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcf, 0xc8, 0x2f,
	0x2e, 0x49, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x47, 0x62, 0x16, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x22, 0x64, 0xf4, 0x10, 0x4c, 0x29, 0x1c, 0x7a,
	0x13, 0x8b, 0x92, 0x33, 0x32, 0xcb, 0x52, 0x21, 0x7a, 0x71, 0x29, 0x4a, 0x49, 0x2c, 0x49, 0x2c,
	0x4e, 0x2d, 0x81, 0x28, 0x52, 0x9a, 0xcc, 0xc8, 0xc5, 0x1e, 0x00, 0xb1, 0x52, 0xc8, 0x8a, 0x8b,
	0x1d, 0x2a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa7, 0x87, 0xd5, 0x7a, 0x3d, 0x17,
	0x88, 0x2a, 0x0f, 0x86, 0x20, 0x98, 0x06, 0x90, 0x5e, 0xa8, 0xed, 0x12, 0x4c, 0x78, 0xf5, 0x3a,
	0x42, 0x54, 0x81, 0xf4, 0x42, 0x35, 0x38, 0x09, 0x70, 0xf1, 0x41, 0x7d, 0x1d, 0x9f, 0x5f, 0x50,
	0x92, 0x99, 0x9f, 0xe7, 0xe4, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0xc6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x39, 0x99, 0x69, 0x25,
	0xa9, 0x29, 0x99, 0x79, 0x99, 0x25, 0xc8, 0xfe, 0xab, 0x40, 0xe6, 0x94, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0xfd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x36, 0x54, 0x34, 0x06, 0x73,
	0x01, 0x00, 0x00,
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PayloadOption != nil {
		{
			size := m.PayloadOption.Size()
			i -= size
			if _, err := m.PayloadOption.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Payload_Dataset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload_Dataset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Dataset != nil {
		{
			size, err := m.Dataset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPayload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Payload_Archive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload_Archive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Archive != nil {
		{
			size, err := m.Archive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPayload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayload(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PayloadOption != nil {
		n += m.PayloadOption.Size()
	}
	return n
}

func (m *Payload_Dataset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dataset != nil {
		l = m.Dataset.Size()
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}
func (m *Payload_Archive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Archive != nil {
		l = m.Archive.Size()
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dataset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Dataset{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.PayloadOption = &Payload_Dataset{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Archive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Archive{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.PayloadOption = &Payload_Archive{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayload
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
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
				return 0, ErrInvalidLengthPayload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayload = fmt.Errorf("proto: unexpected end of group")
)
