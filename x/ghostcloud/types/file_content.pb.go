// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ghostcloud/ghostcloud/file_content.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type FileContent struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *FileContent) Reset()         { *m = FileContent{} }
func (m *FileContent) String() string { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()    {}
func (*FileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01a39ef3605fabc, []int{0}
}
func (m *FileContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileContent.Merge(m, src)
}
func (m *FileContent) XXX_Size() int {
	return m.Size()
}
func (m *FileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileContent.DiscardUnknown(m)
}

var xxx_messageInfo_FileContent proto.InternalMessageInfo

func (m *FileContent) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*FileContent)(nil), "ghostcloud.ghostcloud.FileContent")
}

func init() {
	proto.RegisterFile("ghostcloud/ghostcloud/file_content.proto", fileDescriptor_a01a39ef3605fabc)
}

var fileDescriptor_a01a39ef3605fabc = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0xc8, 0x2f,
	0x2e, 0x49, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x47, 0x62, 0xa6, 0x65, 0xe6, 0xa4, 0xc6, 0x27, 0xe7,
	0xe7, 0x95, 0xa4, 0xe6, 0x95, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x22, 0xa4, 0xf5,
	0x10, 0x4c, 0x25, 0x75, 0x2e, 0x6e, 0xb7, 0xcc, 0x9c, 0x54, 0x67, 0x88, 0x5a, 0x21, 0x09, 0x2e,
	0x76, 0xa8, 0x36, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0xd7, 0xc9, 0xfc, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x64, 0x91, 0x2c, 0xae, 0x40, 0x76, 0x45, 0x49,
	0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0x20, 0x32, 0x09, 0xab, 0x00, 0x00, 0x00,
}

func (m *FileContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintFileContent(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileContent(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileContent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovFileContent(uint64(l))
	}
	return n
}

func sovFileContent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileContent(x uint64) (n int) {
	return sovFileContent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileContent
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
			return fmt.Errorf("proto: FileContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileContent
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
				return ErrInvalidLengthFileContent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFileContent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileContent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileContent
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
func skipFileContent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileContent
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
					return 0, ErrIntOverflowFileContent
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
					return 0, ErrIntOverflowFileContent
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
				return 0, ErrInvalidLengthFileContent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFileContent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFileContent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFileContent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileContent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFileContent = fmt.Errorf("proto: unexpected end of group")
)
